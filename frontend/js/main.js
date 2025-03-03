async function fetchUsers() {
  try {
    const response = await fetch("http://localhost:3000/users");
    if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`);
    const users = await response.json();

    const tbody = document.querySelector("#usersTable tbody");
    tbody.innerHTML = "";

    if (!users || users.length === 0) {
      tbody.innerHTML = '<tr><td colspan="4">No users found</td></tr>';
      return;
    }

    users.forEach((user) => {
      const row = document.createElement("tr");
      row.dataset.userId = user.id;
      row.innerHTML = `
          <td>${user.id}</td>
          <td>${user.name}</td>
          <td>${user.nim}</td>
          <td>${user.email}</td>
        `;
      row.addEventListener("click", () => showUserDetails(user.id));
      tbody.appendChild(row);
    });
  } catch (error) {
    console.error("Error fetching users:", error);
    document
      .getElementById("usersTable")
      .insertAdjacentHTML("afterend", '<p class="error">Error loading users. Please try again later.</p>');
  }
}

async function fetchData(url) {
  try {
    const response = await fetch(url);
    if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`);
    return await response.json();
  } catch (error) {
    console.error("Error fetching data:", error);
    return null;
  }
}

async function loadUsers() {
  try {
    const users = await fetchData("http://localhost:3000/users");
    const tbody = document.querySelector("#usersTable tbody");
    tbody.innerHTML = "";
    if (!users || users.length === 0) {
      tbody.innerHTML = '<tr><td colspan="4">No users found</td></tr>';
      return;
    }
    users.forEach((user) => {
      const row = document.createElement("tr");
      row.dataset.userId = user.id;
      row.innerHTML = `
          <td>${user.id}</td>
          <td>${user.name}</td>
          <td>${user.nim}</td>
          <td>${user.email}</td>
        `;
      row.addEventListener("click", () => showUserDetails(user.id));
      tbody.appendChild(row);
    });
  } catch (error) {
    console.error("Error fetching users:", error);
    document
      .getElementById("usersTable")
      .insertAdjacentHTML("afterend", '<p class="error">Error loading users. Please try again later.</p>');
  }
}

async function showUserDetails(userId) {
  const existingPanel = document.getElementById("userDetailsPanel");
  if (existingPanel) existingPanel.remove();

  const userRow = document.querySelector(`tr[data-user-id="${userId}"]`);
  if (!userRow) {
    console.error("User row not found for ID:", userId);
    return;
  }

  const detailsPanel = document.createElement("div");
  detailsPanel.id = "userDetailsPanel";
  detailsPanel.className = "details-panel";
  detailsPanel.innerHTML = "<p>Loading user details...</p>";
  document.body.appendChild(detailsPanel);

  try {
    const studentCoursesData = await fetchData("http://localhost:3000/studentcourses");
    if (!studentCoursesData || studentCoursesData.length === 0) {
      throw new Error("Could not fetch student courses");
    }

    const studentDetail = studentCoursesData.find((item) => item.user.id === userId);
    if (!studentDetail) {
      throw new Error("Student courses not found for this user");
    }

    let scores = await fetchData(`http://localhost:3000/examscores/${userId}`);
    if (scores === null) scores = [];

    detailsPanel.innerHTML = `
        <h2>User Details</h2>
        <div class="user-info">
          <p><strong>ID:</strong> ${studentDetail.user.id}</p>
          <p><strong>Name:</strong> ${studentDetail.user.name}</p>
          <p><strong>NIM:</strong> ${studentDetail.user.nim}</p>
          <p><strong>Email:</strong> ${studentDetail.user.email}</p>
        </div>
        <h3>Courses Taken</h3>
        <div class="courses-container">
          ${
            studentDetail.courses && studentDetail.courses.length > 0
              ? `<ul class="courses-list">
                ${studentDetail.courses.map((course) => `<li>${course.name} (${course.course_code})</li>`).join("")}
               </ul>`
              : "<p>No courses found for this user.</p>"
          }
        </div>
        <h3>Exam Scores</h3>
        <div class="scores-container">
          ${
            scores.length > 0
              ? `<table class="scores-table">
                 <thead>
                   <tr>
                     <th>Course Code</th>
                     <th>Score</th>
                   </tr>
                 </thead>
                 <tbody>
                   ${scores
                     .map(
                       (score) => `
                     <tr>
                       <td>${score.course_code}</td>
                       <td>${score.score}</td>
                     </tr>
                   `
                     )
                     .join("")}
                 </tbody>
               </table>`
              : "<p>No exam scores found for this user.</p>"
          }
        </div>
        <button id="deleteUserBtn" class="close-btn">Delete User</button>
        <button id="closeDetailsBtn" class="close-btn">Close</button>
      `;

    document.getElementById("closeDetailsBtn").addEventListener("click", () => {
      detailsPanel.remove();
    });

    document.getElementById("deleteUserBtn").addEventListener("click", async () => {
      await deleteUser(userId);
    });
  } catch (error) {
    console.error("Error fetching user details:", error);
    detailsPanel.innerHTML = `
        <p class="error">Error loading details for user ID: ${userId}</p>
        <button id="closeDetailsBtn" class="close-btn">Close</button>
      `;
    document.getElementById("closeDetailsBtn").addEventListener("click", () => {
      detailsPanel.remove();
    });
  }
}

function showNotification(message, type) {
  const notification = document.createElement("div");
  notification.className = `notification ${type}`;
  notification.textContent = message;

  document.body.appendChild(notification);

  setTimeout(() => {
    notification.classList.add("fade-out");
    setTimeout(() => {
      notification.remove();
    }, 500);
  }, 3000);
}

document.addEventListener("DOMContentLoaded", loadUsers);

document.addEventListener('DOMContentLoaded', function() {
    const modal = document.getElementById('addUserModal');
    const addUserBtn = document.getElementById('addUser');
    const closeBtn = document.querySelector('.close');
    const form = document.getElementById('addUserForm');

    addUserBtn.addEventListener('click', function() {
        modal.style.display = 'block';
    });

    closeBtn.addEventListener('click', function() {
        modal.style.display = 'none';
        form.reset();
    });

    window.addEventListener('click', function(event) {
        if (event.target === modal) {
            modal.style.display = 'none';
            form.reset();
        }
    });

    form.addEventListener('submit', async function(event) {
        event.preventDefault();
        
        try {
            await addNewUser();
            modal.style.display = 'none';
            form.reset();
        } catch (error) {
            console.error('Error in form submission:', error);
        }
    });
});

async function addNewUser() {
    const name = document.getElementById('name').value;
    const nim = document.getElementById('nim').value;
    const email = document.getElementById('email').value;

    const userData = {
        name: name,
        nim: nim,
        email: email
    };

    try {
        const response = await fetch('http://localhost:3000/users', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(userData)
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.error || 'Failed to add user');
        }

        const newUser = await response.json();
        showNotification('User added successfully!', 'success');
        fetchUsers();
        
        return newUser;
    } catch (error) {
        showNotification('Error: ' + error.message, 'error');
        throw error;
    }
}

function showNotification(message, type) {
    const notification = document.createElement('div');
    notification.className = `notification ${type}`;
    notification.textContent = message;
    
    document.body.appendChild(notification);
    
    setTimeout(() => {
        notification.classList.add('fade-out');
        setTimeout(() => {
            notification.remove();
        }, 500);
    }, 3000);
}

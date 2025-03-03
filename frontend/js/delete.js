async function deleteUser(userId) {
    if (!confirm('Are you sure you want to delete this user?')) {
      return;
    }
    
    try {
      const response = await fetch(`http://localhost:3000/users/${userId}`, {
        method: 'DELETE',
      });
      
      if (!response.ok) {
        throw new Error(`Failed to delete user: ${response.status}`);
      }
      
      showNotification('User deleted successfully!', 'success');
      
      const detailsPanel = document.getElementById('userDetailsPanel');
      if (detailsPanel) {
        const currentUserElement = detailsPanel.querySelector('.user-info p:first-child');
        if (currentUserElement && currentUserElement.textContent.includes(userId)) {
          detailsPanel.remove();
        }
      }
      
      fetchUsers();
    } catch (error) {
      console.error('Error deleting user:', error);
      showNotification('Error deleting user: ' + error.message, 'error');
    }
  }
  
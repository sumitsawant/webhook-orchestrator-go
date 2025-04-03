document.addEventListener('DOMContentLoaded', function() {
    const form = document.getElementById('reset-password-form');
    const feedback = document.getElementById('feedback');
    const azureAdForm = document.getElementById('azure-ad-form');

    form.addEventListener('submit', function(event) {
        event.preventDefault();
        
        const username = document.getElementById('username').value;
        const oldPassword = document.getElementById('old-password').value;
        const newPassword = document.getElementById('new-password').value;

        if (username && oldPassword && newPassword) {
            // Send a message to the background script to handle the password reset
            chrome.runtime.sendMessage({ action: 'resetPassword', username, oldPassword, newPassword }, function(response) {
                if (response.success) {
                    feedback.textContent = 'Password reset successfully!';
                    feedback.style.color = 'green';
                } else {
                    feedback.textContent = 'Error resetting password: ' + response.error;
                    feedback.style.color = 'red';
                }
            });
        } else {
            feedback.textContent = 'Please fill in all fields.';
            feedback.style.color = 'red';
        }
    });

    azureAdForm.addEventListener('submit', function(event) {
        event.preventDefault();

        const clientId = document.getElementById('client-id').value;
        const clientSecret = document.getElementById('client-secret').value;
        const tenantId = document.getElementById('tenant-id').value;

        if (clientId && clientSecret && tenantId) {
            // Save Azure AD credentials to storage
            chrome.storage.sync.set({ clientId, clientSecret, tenantId }, function() {
                alert('Azure AD credentials saved successfully!');
                // Open reset password page
                document.getElementById('azure-ad-credentials').style.display = 'none';
                document.getElementById('form-container').style.display = 'block';
            });
        } else {
            alert('Please fill in all fields.');
        }
    });

    // Add event listener for the next button
    const nextButton = document.getElementById('next-button');
    nextButton.addEventListener('click', function() {
        document.getElementById('instructions').style.display = 'none';
        document.getElementById('azure-ad-credentials').style.display = 'block';
    });

    // Add event listener for the back button
    document.getElementById('back-button').addEventListener('click', function() {
        document.getElementById('form-container').style.display = 'none';
        document.getElementById('instructions').style.display = 'block';
    });
});
// This file contains the background script for the Chrome extension. 
// It manages events and handles communication between different parts of the extension.

chrome.runtime.onInstalled.addListener(() => {
    console.log("Extension installed");
});

chrome.runtime.onMessage.addListener((request, sender, sendResponse) => {
    if (request.action === "resetPassword") {
        const { username, oldPassword, newPassword } = request;

        // Retrieve Azure AD credentials from storage
        chrome.storage.sync.get(["clientId", "clientSecret", "tenantId"], function(credentials) {
            const { clientId, clientSecret, tenantId } = credentials;

            if (clientId && clientSecret && tenantId) {
                // Azure AD API call to reset password
                fetch(`https://login.microsoftonline.com/${tenantId}/oauth2/v2.0/token`, {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/x-www-form-urlencoded"
                    },
                    body: `client_id=${clientId}&scope=https%3A%2F%2Fgraph.microsoft.com%2F.default&client_secret=${clientSecret}&grant_type=client_credentials`
                })
                .then(response => response.json())
                .then(data => {
                    if (data.access_token) {
                        fetch("https://graph.microsoft.com/v1.0/me/changePassword", {
                            method: "POST",
                            headers: {
                                "Authorization": `Bearer ${data.access_token}`,
                                "Content-Type": "application/json"
                            },
                            body: JSON.stringify({
                                currentPassword: oldPassword,
                                newPassword: newPassword
                            })
                        })
                        .then(response => response.json())
                        .then(data => {
                            if (data.error) {
                                sendResponse({ success: false, error: data.error.message });
                            } else {
                                sendResponse({ success: true, message: "Password reset successfully." });
                            }
                        })
                        .catch(error => {
                            sendResponse({ success: false, error: error.message });
                        });
                    } else {
                        sendResponse({ success: false, error: data.error_description });
                    }
                })
                .catch(error => {
                    sendResponse({ success: false, error: error.message });
                });
            } else {
                sendResponse({ success: false, error: "Azure AD credentials are not set." });
            }
        });

        return true; // Keep the message channel open for sendResponse
    }
});
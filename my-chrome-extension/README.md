# My Chrome Extension

This Chrome extension allows users to reset their Active Directory passwords directly from their browser.

## Project Structure

```
my-chrome-extension
├── src
│   ├── background.js        # Background script for managing events and communication
│   ├── content.js          # Content script for interacting with web pages
│   ├── popup
│   │   ├── popup.html      # HTML structure for the popup interface
│   │   └── popup.js        # JavaScript logic for handling user interactions
│   └── styles
│       └── popup.css       # CSS styles for the popup interface
├── manifest.json           # Configuration file for the Chrome extension
└── README.md               # Documentation for the project
```

## Installation

1. Clone the repository or download the source code.
2. Open Chrome and navigate to `chrome://extensions/`.
3. Enable "Developer mode" by toggling the switch in the top right corner.
4. Click on "Load unpacked" and select the `my-chrome-extension` directory.

## Usage

1. Click on the extension icon in the Chrome toolbar to open the popup.
2. Enter your current password and the new password you wish to set.
3. Click the "Reset Password" button to submit your request.
4. You will receive feedback on the success or failure of the password reset.

## Setting Up Azure AD Connection

To reset the password using the Azure AD API, you need to set up a connection to Azure AD. Follow these steps:

1. Register your application in the Azure portal.
2. Configure the required API permissions for Microsoft Graph.
3. Obtain the client ID and client secret for your application.
4. Use the OAuth 2.0 client credentials flow to obtain an access token. You can use tools like Postman or write a script to get the token.

Example request to obtain the access token:

```
POST https://login.microsoftonline.com/{tenant}/oauth2/v2.0/token
Content-Type: application/x-www-form-urlencoded

client_id={client_id}&scope=https%3A%2F%2Fgraph.microsoft.com%2F.default&client_secret={client_secret}&grant_type=client_credentials
```

Replace `{tenant}`, `{client_id}`, and `{client_secret}` with your actual tenant ID, client ID, and client secret.

## Contributing

Feel free to submit issues or pull requests for improvements or bug fixes. 

## License

This project is licensed under the MIT License.
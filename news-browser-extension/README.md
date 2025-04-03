# News Browser Extension

This browser extension displays live news directly in your browser. It fetches the latest news articles and presents them in a user-friendly popup interface.

## Features

- Live news updates from various sources.
- Popup interface for easy access to news articles.
- Customizable settings for news sources.

## Project Structure

```
news-browser-extension
├── src
│   ├── background.js        # Background script for managing events and communication
│   ├── content.js          # Content script for fetching and displaying news on web pages
│   ├── popup
│   │   ├── popup.html      # HTML structure for the popup
│   │   ├── popup.js        # JavaScript for handling popup interactions
│   │   └── popup.css       # Styles for the popup
│   └── manifest.json       # Configuration file for the browser extension
├── package.json            # npm configuration file
└── README.md               # Documentation for the project
```

## Installation

1. Clone the repository:
   ```
   git clone <repository-url>
   ```
2. Navigate to the project directory:
   ```
   cd news-browser-extension
   ```
3. Install dependencies:
   ```
   npm install
   ```

## Usage

1. Load the extension in your browser:
   - Go to `chrome://extensions/` (for Chrome) or `about:debugging` (for Firefox).
   - Enable "Developer mode" (for Chrome).
   - Click on "Load unpacked" and select the `src` directory.
2. Click on the extension icon to open the popup and view live news.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
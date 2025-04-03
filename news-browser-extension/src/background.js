const newsAPI = 'https://newsapi.org/v2/top-headlines?country=us&apiKey=YOUR_API_KEY';

chrome.runtime.onInstalled.addListener(() => {
    console.log('News Browser Extension installed');
});

chrome.runtime.onMessage.addListener((request, sender, sendResponse) => {
    if (request.action === 'fetchNews') {
        fetch(newsAPI)
            .then(response => response.json())
            .then(data => {
                sendResponse({ news: data.articles });
            })
            .catch(error => {
                console.error('Error fetching news:', error);
                sendResponse({ news: [] });
            });
        return true; // Indicates that the response will be sent asynchronously
    }
});
const newsApiUrl = 'https://newsapi.org/v2/top-headlines?country=us&apiKey=YOUR_API_KEY';

async function fetchNews() {
    try {
        const response = await fetch(newsApiUrl);
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        const data = await response.json();
        displayNews(data.articles);
    } catch (error) {
        console.error('Error fetching news:', error);
    }
}

function displayNews(articles) {
    const newsContainer = document.createElement('div');
    newsContainer.style.position = 'fixed';
    newsContainer.style.top = '10px';
    newsContainer.style.right = '10px';
    newsContainer.style.backgroundColor = 'white';
    newsContainer.style.border = '1px solid #ccc';
    newsContainer.style.padding = '10px';
    newsContainer.style.zIndex = '1000';

    articles.forEach(article => {
        const articleElement = document.createElement('div');
        articleElement.innerHTML = `<h4>${article.title}</h4><p>${article.description}</p><a href="${article.url}" target="_blank">Read more</a>`;
        newsContainer.appendChild(articleElement);
    });

    document.body.appendChild(newsContainer);
}

fetchNews();
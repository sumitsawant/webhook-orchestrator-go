const newsApiUrl = 'https://newsapi.org/v2/top-headlines?country=us&apiKey=YOUR_API_KEY';

document.addEventListener('DOMContentLoaded', function() {
    const newsContainer = document.getElementById('news-container');
    const refreshButton = document.getElementById('refresh-button');
    // Removed search bar reference

    function fetchNews() {
        fetch(newsApiUrl)
            .then(response => response.json())
            .then(data => {
                displayNews(data.articles);
            })
            .catch(error => {
                console.error('Error fetching news:', error);
                newsContainer.innerHTML = '<p>Error fetching news. Please try again later.</p>';
            });
    }

    function displayNews(articles) {
        newsContainer.innerHTML = '';
        articles.forEach(article => {
            const articleElement = document.createElement('div');
            articleElement.classList.add('news-article');
            articleElement.innerHTML = `
                <h3>${article.title}</h3>
                <p>${article.description}</p>
                <a href="${article.url}" target="_blank">Read more</a>
            `;
            newsContainer.appendChild(articleElement);
        });
    }

    // Removed filterNews function

    // Removed search bar event listener

    refreshButton.addEventListener('click', fetchNews);
    fetchNews();
});
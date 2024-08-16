function searchArtists(query) {
    if (query.length === 0) {
        document.getElementById('search-results').innerHTML = '';
        return;
    }

    fetch(`/search?q=${query}`)
        .then(response => response.json())
        .then(data => {
            let resultsContainer = document.getElementById('search-results');
            resultsContainer.innerHTML = '';

            data.forEach(result => {
                let resultItem = document.createElement('div');
                resultItem.className = 'search-result-item';
                resultItem.innerHTML = result.name;

                resultItem.onclick = function() {
                    // Redirect based on the type of result
                    if (result.name.includes(" - member of ")) {
                        let artistName = result.name.split(" - member of ")[1];
                        fetch(`/search?q=${encodeURIComponent(artistName)}`)
                            .then(response => response.json())
                            .then(artistData => {
                                if (artistData.length > 0) {
                                    window.location.href = `/artist/${artistData[0].id}`;
                                }
                            });
                    } else {
                        window.location.href = `/artist/${result.id}`;
                    }
                };
                resultsContainer.appendChild(resultItem);
            });
        })
        .catch(error => console.error('Error:', error));
}

// Add a script to handle the toggle functionality for the search bar and navigation links.
document.addEventListener("DOMContentLoaded", function () {
    const hamburger = document.querySelector(".hamburger");
    const navLinks = document.querySelector(".nav-links");
    const searchIcon = document.querySelector(".search-icon");
    const search = document.querySelector(".search");

    hamburger.addEventListener("click", () => {
        navLinks.classList.toggle("active");
    });

    searchIcon.addEventListener("click", () => {
        search.classList.toggle("active");
    });
});

// Function to search for artists based on user query
function searchArtists(query) {
    console.log("Search query:", query);
    const resultsContainer = document.getElementById('search-results');

    if (query.length === 0) {
        resultsContainer.innerHTML = '';
        return;
    }

    fetch(`/search?q=${encodeURIComponent(query)}`)
        .then(response => response.json())
        .then(data => {
            resultsContainer.innerHTML = '';

            // Display each search result
            data.forEach(result => {
                const resultItem = document.createElement('div');
                resultItem.className = 'search-result-item';
                resultItem.textContent = result.name;

                // Direct the user to the artist detail page
                resultItem.onclick = function() {
                    window.location.href = `/artist/${result.id}`;
                };

                resultsContainer.appendChild(resultItem);
            });
        })
        .catch(error => console.error('Error fetching search results:', error));
}

// Initialize map for artist details page
document.addEventListener("DOMContentLoaded", function() {
    const mapContainer = document.getElementById('map');
    if (mapContainer) {
        initializeMapbox();
    }
});

function initializeMapbox() {
    mapboxgl.accessToken = 'pk.eyJ1IjoiYXRob29oIiwiYSI6ImNtMWY2N3prZjJsN3MybHNjMWd3bThzOXcifQ.HNgAHQBkzGdrnuS1MtwYlQ';

    const concertLocations = JSON.parse(document.getElementById('ConcertLocationsJSON').innerText);

    const map = new mapboxgl.Map({
        container: 'map',
        style: 'mapbox://styles/mapbox/streets-v11',
        center: [0, 0],
        zoom: 2 // Default zoom level
    });

    concertLocations.forEach(function(location) {
        const popup = new mapboxgl.Popup({ offset: 25 }).setText(location.locationName);

        new mapboxgl.Marker()
            .setLngLat(location.coordinates)
            .setPopup(popup)
            .addTo(map);
    });

    // Ensure that the map resizes correctly
    map.resize();
}


// Toggle functionality for search bar and navigation links
document.addEventListener("DOMContentLoaded", function() {
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


// Optional: Additional configuration if you want to control the slider behavior.
document.addEventListener('DOMContentLoaded', function() {
    const slideTrack = document.querySelector('.slide-track');
    const slideItems = document.querySelectorAll('.slide-item');

    // Clone the slide items to make the slider continuous
    slideItems.forEach(item => {
        slideTrack.appendChild(item.cloneNode(true));
    });
});



// header scrolls

document.addEventListener('DOMContentLoaded', function () {
    const header = document.querySelector('header');
    const hamburger = document.querySelector('.hamburger');
    const navLinks = document.querySelector('.nav-links');
    let lastScrollY = window.scrollY;

    // Hide/Show Header on Scroll
    window.addEventListener('scroll', () => {
        if (window.scrollY > lastScrollY) {
            header.style.top = '-100px';  // Hide header on scroll down
        } else {
            header.style.top = '0';       // Show header on scroll up
        }
        lastScrollY = window.scrollY;
    });
});

// Hamburger menu functionality
document.addEventListener('DOMContentLoaded', function () {
    const hamburger = document.querySelector('.hamburger');
    const navLinks = document.querySelector('.nav-links');

    hamburger.addEventListener('click', () => {
        navLinks.classList.toggle('active');
        hamburger.classList.toggle('active'); // Toggles the "X" icon
    });
});

document.addEventListener('DOMContentLoaded', function() {
    fetch('/api/artists')
        .then(response => response.json())
        .then(data => {
            const artistsDiv = document.getElementById('artists');
            data.forEach(artist => {
                const artistDiv = document.createElement('div');
                artistDiv.innerHTML = `<h2>${artist.name}</h2><p>${artist.firstAlbum}</p>`;
                artistsDiv.appendChild(artistDiv);
            });
        })
        .catch(error => console.error('Error fetching artists:', error));
});

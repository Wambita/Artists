//search bar functionality
let searchTimeout;

document.getElementById("search-bar").addEventListener("input", function () {
    clearTimeout(searchTimeout);

    const query = this.value.trim();
    if (query.length > 0) {
        searchTimeout = setTimeout(() => {
            fetchSuggestions(query);
        }, 300); // Adjust debounce time as needed
    }
});

async function fetchSuggestions(query) {
    try {
        const response = await fetch(`/search?q=${encodeURIComponent(query)}`);
        if (!response.ok) throw new Error("Failed to fetch suggestions");

        const suggestions = await response.json();
        displaySuggestions(suggestions);
    } catch (error) {
        console.error("Error fetching suggestions:", error);
    }
}

function displaySuggestions(suggestions) {
    const suggestionBox = document.getElementById("suggestions");
    suggestionBox.innerHTML = ""; // Clear previous suggestions

    suggestions.forEach(suggestion => {
        const item = document.createElement("div");
        item.classList.add("suggestion-item");
        item.innerText = `${suggestion.Name} - ${suggestion.Type}`;
        item.addEventListener("click", () => selectSuggestion(suggestion.ID));
        suggestionBox.appendChild(item);
    });
}

function selectSuggestion(id) {
    // Use the artist ID to redirect to the artist's page or handle selection
    window.location.href = `/artist?id=${id}`;
}


//drop down menu for locations dates
function toggleViews() {
    const viewSelect = document.getElementById("view-select");
    const locationsSection = document.getElementById("locations");
    const datesSection = document.getElementById("dates");
    const locationsAndDatesSection = document.getElementById("locationsAndDates");

    // Hide all sections by default
    locationsSection.classList.add("hidden");
    datesSection.classList.add("hidden");
    locationsAndDatesSection.style.display = "none";

    // Show the selected section
    switch (viewSelect.value) {
        case "locations":
            locationsSection.classList.remove("hidden");
            break;
        case "dates":
            datesSection.classList.remove("hidden");
            break;
        case "locationsAndDates":
            locationsAndDatesSection.style.display = "block";
            break;
    }
}

// Initially display the default view (Locations and Dates)
toggleViews();
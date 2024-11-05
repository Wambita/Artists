//search bar functionality
let searchTimeout;

document.getElementById("search-bar").addEventListener("input", function () {
    clearTimeout(searchTimeout);

    const query = this.value.trim();
    if (query.length > 0) {
        searchTimeout = setTimeout(() => {
            fetchSuggestions(query);
        }, 300); // Adjust debounce time as needed
    }else{
        document.getElementById("suggestions").innerHTML = ""
        document.getElementById("suggestions").classList = "hidden"
    }
});

async function fetchSuggestions(query) {
    try {
        const response = await fetch(`/search?q=${encodeURIComponent(query)}`);
        if (!response.ok) throw new Error("Failed to fetch suggestions");

        const suggestions = await response.json();
        // console.log(suggestions)

        displaySuggestions(suggestions);
    } catch (error) {
        console.error("Error fetching suggestions:", error);
    }
}

function displaySuggestions(suggestions) {
    const suggestionBox = document.getElementById("suggestions");
    suggestionBox.innerHTML = ""; // Clear previous suggestions
    
    if (suggestions != null){
        suggestionBox.classList = "suggestion-container";

        suggestions.forEach(suggestion => {
            const item = document.createElement("li");
            item.innerText = `${suggestion.name}   (${suggestion.type})`;
            item.addEventListener("click", () => selectSuggestion(suggestion.id));
            suggestionBox.appendChild(item);
        });
    }else{
        suggestionBox.classList = "hidden";
    }  
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
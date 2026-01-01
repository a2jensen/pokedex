const API_BASE = 'http://localhost:8080';

let prevUrl = null;
let nextUrl = null;

// DOM elements
const locationList = document.getElementById('location-list');
const pokemonList = document.getElementById('pokemon-list');
const caughtList = document.getElementById('caught-list');
const currentLocation = document.getElementById('current-location');
const prevBtn = document.getElementById('prev-btn');
const nextBtn = document.getElementById('next-btn');
const messageEl = document.getElementById('message');

// Show message
function showMessage(text, isError = false) {
    messageEl.textContent = text;
    messageEl.className = isError ? 'show error' : 'show';
    setTimeout(() => {
        messageEl.className = '';
    }, 3000);
}

// Fetch locations
async function fetchLocations(url = null) {
    try {
        const endpoint = url || `${API_BASE}/map`;
        const res = await fetch(endpoint);
        const data = await res.json();

        prevUrl = data.previous;
        nextUrl = data.next;

        prevBtn.disabled = !prevUrl;
        nextBtn.disabled = !nextUrl;

        locationList.innerHTML = '';
        data.results.forEach(loc => {
            const div = document.createElement('div');
            div.className = 'location-item';
            div.textContent = loc.name;
            div.onclick = () => exploreLocation(loc.name);
            locationList.appendChild(div);
        });
    } catch (err) {
        showMessage('Failed to fetch locations', true);
    }
}

// Explore a location
async function exploreLocation(name) {
    try {
        currentLocation.textContent = `Exploring: ${name}`;
        const res = await fetch(`${API_BASE}/explore/${name}`);
        const data = await res.json();

        pokemonList.innerHTML = '';
        data.pokemon.forEach(p => {
            const div = document.createElement('div');
            div.className = 'pokemon-item';
            div.textContent = p.name;
            div.onclick = () => catchPokemon(p.name);
            pokemonList.appendChild(div);
        });
    } catch (err) {
        showMessage('Failed to explore location', true);
    }
}

// Catch a pokemon
async function catchPokemon(name) {
    try {
        const res = await fetch(`${API_BASE}/catch/${name}`, { method: 'POST' });
        const data = await res.json();

        if (data.caught) {
            showMessage(`Caught ${name}!`);
            fetchPokedex();
        } else {
            showMessage(`${name} escaped!`, true);
        }
    } catch (err) {
        showMessage('Failed to catch pokemon', true);
    }
}

// Fetch pokedex
async function fetchPokedex() {
    try {
        const res = await fetch(`${API_BASE}/pokedex`);
        const data = await res.json();

        caughtList.innerHTML = '';
        data.pokemon.forEach(p => {
            const div = document.createElement('div');
            div.className = 'caught-item';
            div.textContent = p.name;
            caughtList.appendChild(div);
        });
    } catch (err) {
        showMessage('Failed to fetch pokedex', true);
    }
}

// Event listeners
prevBtn.onclick = () => fetchLocations(prevUrl);
nextBtn.onclick = () => fetchLocations(nextUrl);

// Init
fetchLocations();
fetchPokedex();

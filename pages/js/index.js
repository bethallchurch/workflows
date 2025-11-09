(async () => {
    const response = await fetch('../data/w3wAutoSuggestions.json');
    const data = await response.json();

    const el = document.querySelector('.data');
    el.innerHTML = JSON.stringify(data, null, 4);
})();

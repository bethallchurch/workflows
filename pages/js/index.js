import data from '../data/w3wAutoSuggestions.json' with { type: 'json' };

const el = document.querySelector('.data');
el.innerHTML = JSON.stringify(data, null, 4);

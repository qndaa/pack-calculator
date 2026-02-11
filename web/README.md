# Pack Calculator UI

A simple web interface to calculate how many packs you need for a given number of items.

## What it does

Enter the number of items you want to ship, and the calculator figures out the optimal combination of pack sizes. No more guessing!

## How to use

1. Start the backend server (see main README)
2. Open http://localhost:8080 in your browser
3. You'll see the available pack sizes at the top
4. Type in how many items you need
5. Hit "Calculate" and you're done

## Tech stack

Nothing fancy here — just plain HTML, CSS, and vanilla JavaScript. No build tools, no npm, no frameworks. It just works.

## Files

- `index.html` — that's it, everything lives in one file

## API endpoints used

- `GET /packs` — fetches available pack sizes on page load
- `POST /calculate` — sends the item count and gets back the pack breakdown

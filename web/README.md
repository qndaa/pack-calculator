# Pack Calculator UI

A simple web interface to calculate how many packs you need for a given number of items.

## What it does

Enter the number of items you want to ship, and the calculator figures out the optimal combination of pack sizes. It follows three rules:

1. **Only whole packs** — packs can't be broken open
2. **Minimal overship** — send the least amount of items possible
3. **Fewest packs** — when tied, use fewer boxes

## Features

- View available pack sizes
- Add new pack sizes
- Delete existing pack sizes  
- Calculate optimal pack combination
- See total items and pack count

## How to use

1. Start the backend server (see main README)
2. Open http://localhost:8080 in your browser
3. You'll see the available pack sizes at the top
4. Add or remove pack sizes as needed
5. Type in how many items you need
6. Hit "Calculate" and you're done

## Tech stack

Nothing fancy here — just plain HTML, CSS, and vanilla JavaScript. No build tools, no npm, no frameworks. It just works.

## Files

- `index.html` — that's it, everything lives in one file

## API endpoints used

- `GET /packs` — fetches available pack sizes on page load
- `POST /packs/{size}` — adds a new pack size
- `DELETE /packs/{size}` — removes a pack size
- `POST /calculate` — sends the item count and gets back the pack breakdown

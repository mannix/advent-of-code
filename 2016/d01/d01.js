const fs = require('fs');

// Starting position
let position = {
  x: 0,
  y: 0
};
// Starting direction
let direction = 'N';
// List of every coordinate we step on
let coordinates = [];
// Add our starting location
coordinates.push({...position});

const movesFromInputFile = () => {
  try {
    const data = fs.readFileSync('input', 'utf8');
    return data.split(", ");
  } catch (error) {
    console.log(`Error reading input file: \n${error}`);
    process.exit(1);
  }
};

const moveHorizontally = (amount) => {
  position.x += amount;
  coordinates.push({...position});
}

const moveVertically = (amount) => {
  position.y += amount;
  coordinates.push({...position});
}

const processMove = (move) => {
  // Turn direction for this move
  const turn = move.substring(0, 1);
  const steps = parseInt(move.substring(1, move.length), 10);

  // Update direction based on turn
  switch (direction) {
    case 'N':
      direction = turn == 'L' ? 'W' : 'E';
      break;
    case 'E':
      direction = turn == 'L' ? 'N' : 'S';
      break;
    case 'S':
      direction = turn == 'L' ? 'E' : 'W';
      break;
    case 'W':
      direction = turn == 'L' ? 'S' : 'N';
      break;
  }

  // Change position based on new direction
  switch (direction) {
    case 'N':
      for (i = 1; i <= steps; i++) {
        moveVertically(1);
      }
      break;
    case 'E':
      for (i = 1; i <= steps; i++) {
        moveHorizontally(1);
      }
      break;
    case 'S':
      for (i = 1; i <= steps; i++) {
        moveVertically(-1);
      }
      break;
    case 'W':
      for (i = 1; i <= steps; i++) {
        moveHorizontally(-1);
      }
      break;
  }
}

console.log("\n----------------------------------------");

// Get array of moves from input
const moves = movesFromInputFile();

// Loop through steps updating position and direction each time
moves.forEach((move) => {
  processMove(move);
});

console.log(`\nPart 1 distance = ${Math.abs(position.x) + Math.abs(position.y)}`);

coordinates.forEach((coordinate, index) => {
  const leftCoordinates = coordinates.slice(0, index);
  if (leftCoordinates.some(c => c.x == coordinate.x && c.y == coordinate.y)) {
    console.log(`\nPart 2 distance = ${Math.abs(coordinate.x) + Math.abs(coordinate.y)}`);
    console.log("\n----------------------------------------");
    process.exit(0);
  }
});


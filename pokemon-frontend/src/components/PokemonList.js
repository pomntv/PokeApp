// import React, { useState, useEffect } from 'react';

// function PokemonList() {
//   const [pokemons, setPokemons] = useState([]);
//   const [loading, setLoading] = useState(true);

//   useEffect(() => {
//     // fetch('http://localhost:8080/api/pokemon/all')
//     fetch('http://localhost:8080/api/pokemon')

//       .then((response) => response.json())
//       .then((data) => {
//         setPokemons(data);
//         setLoading(false);
//       })
//       .catch((error) => {
//         console.error("Error fetching data:", error);
//         setLoading(false);
//       });
//   }, []);

//   if (loading) {
//     return <div>Loading...</div>;
//   }

//   return (
//     <div>
//       <h1>Pokémon List</h1>
//       <ul>
//         {pokemons.map((pokemon) => (
//           <li key={pokemon.ID}>
//             <img 
//               src={`https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/${pokemon.ID}.png`} 
//               alt={pokemon.Name} 
//             />
//             {pokemon.Name}
//           </li>
//         ))}
//       </ul>
//     </div>
//   );
// }

// export default PokemonList;

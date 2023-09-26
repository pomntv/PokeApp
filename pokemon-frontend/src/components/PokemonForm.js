import React, { useState, useEffect } from 'react';
import axios from 'axios';

function PokemonList() {
    const [pokemons, setPokemons] = useState([]);

    useEffect(() => {
        axios.get('http://localhost:8080/api/pokemon/all')
            .then(response => {
                setPokemons(response.data);
            })
            .catch(error => {
                console.error("Error fetching data:", error);
            });
    }, []);

    return (
        <div>
            <h2>All Pokemon</h2>
            <ul>
                {pokemons.map(pokemon => (
                    <li key={pokemon._id}>{pokemon.name}</li>
                ))}
            </ul>
        </div>
    );
}

export default PokemonList;

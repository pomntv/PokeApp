import React, { useState, useEffect } from 'react';
import axios from 'axios';
import Select from 'react-select';

function PokemonForm() {
  const [team, setTeam] = useState([{name: '', id: ''}, {name: '', id: ''}, {name: '', id: ''}, {name: '', id: ''}, {name: '', id: ''}, {name: '', id: ''}]);
  const [pokemonOptions, setPokemonOptions] = useState([]);

  useEffect(() => {
    axios.get('http://localhost:8080/api/pokemon')
      .then(response => {
        const options = response.data.map(pokemon => ({ value: pokemon.name, label: pokemon.name, id: pokemon.id }));
        setPokemonOptions(options);
      })
      .catch(error => {
        console.error("Error fetching data:", error);
      });
  }, []);

  const handleSelectChange = (selectedOption, index) => {
    const newTeam = [...team];
    newTeam[index] = {name: selectedOption.value, id: selectedOption.id};
    setTeam(newTeam);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    // Handle the submission logic here
  };

  const pokemonTypes = ['Normal', 'Fire', 'Water', 'Grass', 'Electric', 'Ice', 'Fighting', 'Poison', 'Ground', 'Flying'];

  return (
    <div>
      <h1>Pokémon Team Analysis</h1>
      <form onSubmit={handleSubmit}>
        {team.map((pokemon, index) => (
          <div key={index}>
            <label>
              Pokémon {index + 1}:
              <Select 
                value={{ value: pokemon.name, label: pokemon.name }}
                onChange={(selectedOption) => handleSelectChange(selectedOption, index)}
                options={pokemonOptions}
                isSearchable={true}
                placeholder="Select a Pokémon"
              />
            </label>
          </div>
        ))}
        <button type="submit">Analyze</button>
      </form>

      <table>
        <thead>
          <tr>
            <th>Attack Type</th>
            {team.map(pokemon => (
              <th key={pokemon.name}>
                <img src={`https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/${pokemon.id}.png`} alt={pokemon.name} />
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {pokemonTypes.map(type => (
            <tr key={type}>
              <td>{type}</td>
              {team.map(pokemon => (
                <td key={pokemon.name}> 
                   
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default PokemonForm;

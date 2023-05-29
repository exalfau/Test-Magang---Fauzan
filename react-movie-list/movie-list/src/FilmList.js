import React, { useState } from 'react';
import filmsData from './FilmData'; // Data film yang akan ditampilkan

const FilmList = () => {
  const [films, setFilms] = useState(filmsData);
  const [searchTerm, setSearchTerm] = useState('');
  const [sortOrder, setSortOrder] = useState('asc');

  // Fungsi untuk mengubah nilai pencarian
  const handleSearch = (event) => {
    setSearchTerm(event.target.value);
  };

  // Fungsi untuk mengubah urutan pengurutan
  const handleSort = (event) => {
    setSortOrder(event.target.value);
  };

  // Fungsi untuk mengurutkan film berdasarkan durasi
  const sortFilms = (films) => {
    if (sortOrder === 'asc') {
      return films.sort((a, b) => a.duration - b.duration);
    } else {
      return films.sort((a, b) => b.duration - a.duration);
    }
  };

  // Fungsi untuk mencari film berdasarkan judul
  const searchFilms = (films) => {
    return films.filter((film) =>
      film.title.toLowerCase().includes(searchTerm.toLowerCase())
    );
  };

  // Mengurutkan film dan menerapkan pencarian
  const sortedFilms = sortFilms(films);
  const searchedFilms = searchFilms(sortedFilms);

  return (
    <div>
      <div className="search-container">
        <input
          type="text"
          placeholder="Cari film berdasarkan judul"
          value={searchTerm}
          onChange={handleSearch}
        />
        <select value={sortOrder} onChange={handleSort}>
          <option value="asc">Urutkan berdasarkan durasi (asc)</option>
          <option value="desc">Urutkan berdasarkan durasi (desc)</option>
        </select>
      </div>
      <div className="film-list">
        {searchedFilms.map((film) => (
          <div key={film.id} className="film-item">
            <img src={film.image} alt={film.title} />
            <h3>{film.title}</h3>
            <p>{film.description}</p>
            <p>Duration: {film.duration} minutes</p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default FilmList;

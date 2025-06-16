// reuseable card component for displaying album information on the home / library page
import "../assets/styles/layout.css";
import "../assets/styles/card.css";
import React from 'react';

type AlbumCardProps = {
  albumName: string;
  artistName?: string;
  coverPath: string; // relative or absolute path to the cover image
  onClick?: () => void;
};

const AlbumCard: React.FC<AlbumCardProps> = ({ albumName, artistName, coverPath, onClick }) => {
  return (
    <div className="album-card" onClick={onClick}>
      <div className="album-cover">
        <img src={coverPath} alt={`${albumName} cover`} loading="lazy" />
      </div>
      <div className="album-info">
        <h3 className="album-title">{albumName}</h3>
        <p className="album-artist">{artistName || 'Unknown Artist'}</p>
      </div>
    </div>
  );
};

export default AlbumCard;


import { useNavigate } from 'react-router-dom';
import "../assets/styles/card.css";

type AlbumCardProps = {
  albumName: string;
  artistName?: string;
  coverPath: string;
};

const AlbumCard = ({ albumName, artistName, coverPath }: AlbumCardProps) => {
  const navigate = useNavigate();

  return (
    <div className="album-card" onClick={() => navigate(`/album/${encodeURIComponent(albumName)}`)}>
      <div className="album-cover">
      <img src={`file:///${coverPath}`} alt="Missing Album Cover" />
      </div>
      <div className="album-info">
        <h3 className="album-title">{albumName}</h3>
        <p className="album-artist">{artistName || 'Unknown Artist'}</p>
      </div>
    </div>
  );
};

export default AlbumCard;


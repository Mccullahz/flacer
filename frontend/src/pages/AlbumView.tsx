import { useNavigate, useParams } from 'react-router-dom';
import { useEffect, useState } from 'react';
import { GetAllTracks } from '../../wailsjs/go/libmanager/Service';
import ManageMenu from '../components/ManageMenu';
import '../assets/styles/card.css';

type Track = Awaited<ReturnType<typeof GetAllTracks>>[number];

const AlbumPage = () => {
  const navigate = useNavigate();
  const { albumName } = useParams();
  const [tracks, setTracks] = useState<Track[]>([]);
  const [artist, setArtist] = useState<string>('Unknown Artist');
  const [coverPath, setCoverPath] = useState<string>('');

  useEffect(() => {
    const fetch = async () => {
      const allTracks = await GetAllTracks();
      const albumTracks = allTracks.filter(track => track.album === albumName);

      if (albumTracks.length > 0) {
        setArtist(albumTracks[0].artist || 'Unknown Artist');
	setCoverPath(albumTracks[0].coverPath || '');
      }

      setTracks(albumTracks);
    };
    fetch();
  }, [albumName]);

const [showManageMenu, setShowManageMenu] = useState(false);

const handleDeleteAlbum = () => {
  alert("Album deleted!"); {/* TODO: hook into Go*/}
  setShowManageMenu(false);
};

const toggleManageMenu = () => setShowManageMenu(!showManageMenu);


  return (
    <div className="album-view-container">
      <button onClick={() => navigate(-1)} className="back-button">
        ⬅ Back
      </button>
      <button onClick={(toggleManageMenu)} className="manage-button">
      ⚙ 
      </button>
	{showManageMenu && (
	  <ManageMenu
	    onClose={() => setShowManageMenu(false)}
	    onDelete={handleDeleteAlbum}
	  />
	)}
      <div className="album-header">
        <img src={coverPath} alt="Missing Album Cover" className="album-header-cover" />
        <div className="album-header-info">
          <h1>{albumName}</h1>
          <p>{artist}</p>
        </div>
      </div>

	<div className="album-header-actions">
 	     <button onClick={() => {}} className="play-all-button">
	     ▶ Play
	      </button>
      	</div>

      <ul className="track-list">
        {tracks.map((track) => (
          <li key={track.id} className="track-item">
            <span className="track-play">▶</span>
            <span className="track-title">{track.title}</span>
            <span className="track-duration">--:--</span> {/* duration from json not ready to go yet */ }
          </li>
        ))}
      </ul>
    </div>
  );
};

export default AlbumPage;

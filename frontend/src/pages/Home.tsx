import { useEffect, useState } from 'react';
import AlbumCard from '../components/AlbumCard';
import '../assets/styles/card.css';

type Track = {
  id: string;
  filePath: string;
  title: string;
  format: string;
  album: string;
  original: string;
  dateAdded: string;
};

type AlbumInfo = {
  albumName: string;
  artistName?: string;
  coverPath: string;
};

const Home = () => {
  const [albums, setAlbums] = useState<AlbumInfo[]>([]);

  useEffect(() => {
    const loadLibrary = async () => {
      try {
        const tracks: Track[] = await window.libmanager.GetAllTracks(); // not currently implemented, but in progress
        const albumMap = new Map<string, AlbumInfo>();

        for (const track of tracks) {
          const album = track.album || 'Unknown Album';
          if (!albumMap.has(album)) {
            const basePath = track.filePath.substring(0, track.filePath.lastIndexOf('/'));
            const coverCandidates = ['cover.jpg', 'cover.png', 'folder.jpg', 'folder.png'];
            let coverPath = '';

            for (const candidate of coverCandidates) {
              const fullPath = `${basePath}/${candidate}`;
              // Optionally, check if file exists via API; here we just assume
              coverPath = fullPath;
              break;
            }

            albumMap.set(album, {
              albumName: album,
              artistName: 'Unknown Artist', // add artist support later
              coverPath,
            });
          }
        }

        setAlbums(Array.from(albumMap.values()));
      } catch (error) {
        console.error('Failed to load library:', error);
      }
    };

    loadLibrary();
  }, []);

  return (
    <div className="album-grid-container">
      {albums.length === 0 ? (
        <p style={{ color: 'white' }}>No albums found.</p>
      ) : (
        albums.map((album, index) => (
          <AlbumCard
            key={index}
            albumName={album.albumName}
            artistName={album.artistName}
            coverPath={album.coverPath}
          />
        ))
      )}
    </div>
  );
};

export default Home;


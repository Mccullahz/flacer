import '../assets/styles/layout.css';
import { AddMusicFolder, OpenDirectorySelector, RescanLibrary } from '../../wailsjs/go/libmanager/Service';

export default function Sidebar() {
const handleAddMusic = async () => {
  try {
    const folder = await OpenDirectorySelector();
    if (!folder) return;

    const tracks = await AddMusicFolder(folder);
    alert(`Added ${tracks.length} tracks to library.`);
  } catch (error) {
    console.error("Error adding music folder:", error);
    alert("Failed to add music from folder.");
  }
  RescanLibrary(); // refresh the library after adding new music
};

  return (
    <aside className="sidebar">
      <div className="logo">LOGO</div>
      <button className="add-button" onClick={handleAddMusic}>
        + Add Music
      </button>
    </aside>
  );
}


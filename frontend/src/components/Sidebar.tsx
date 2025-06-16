import '../assets/styles/layout.css';
import { AddMusicFile, OpenFileSelector } from '../../wailsjs/go/libmanager/Service';

export default function Sidebar() {
  const handleAddMusic = async () => {
    try {
      // Get the context to access the Wails runtime
      const filePath = await OpenFileSelector();
      if (!filePath) return;

      await AddMusicFile(filePath);
      alert(`Added ${filePath}`);
    } catch (error) {
      console.error("Error adding music:", error);
      alert("Failed to add music. Please try again.");
    }
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


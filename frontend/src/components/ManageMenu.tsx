import { useEffect, useRef } from 'react';

interface ManageMenuProps {
  onClose: () => void;
  onDelete: () => void;
  onEdit?: () => void; // future use
}

export default function ManageMenu({ onClose, onDelete, onEdit }: ManageMenuProps) {
  const menuRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handleClickOutside = (e: MouseEvent) => {
      if (menuRef.current && !menuRef.current.contains(e.target as Node)) {
        onClose();
      }
    };

    document.addEventListener('mousedown', handleClickOutside);
    return () => document.removeEventListener('mousedown', handleClickOutside);
  }, [onClose]);

  return (
    <div className="manage-popup" ref={menuRef}>
      <button onClick={onDelete}>Delete Album</button>
      {onEdit && <button onClick={onEdit}>Edit Metadata</button>}
      <button onClick={onClose}>Close</button>
    </div>
  );
}


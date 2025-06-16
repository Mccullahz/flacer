// this file is the core shell component for the app layout
import Header from './Header';
import Sidebar from './Sidebar';
import { ReactNode } from 'react';
import "../assets/styles/layout.css";


export default function Layout({ children }: { children: ReactNode }) {
  return (
    <>
      <Header />
      <Sidebar />
      <main className="main-content">
        {children}
      </main>
    </>
  );
}


import React from 'react';
import Navbar from '../components/Navbar';

function Home() {
  return (
    <div>
      <Navbar />
      <main className="p-4">
        <h1 className="text-2xl font-bold">Welcome to My Portfolio</h1>
      </main>
    </div>
  );
}
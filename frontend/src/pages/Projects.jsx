import React, { useState, useEffect } from 'react';
import Navbar from '../components/Navbar';
import ProjectGrid from '../components/ProjectGrid';

function Projects() {
  const [projects, setProjects] = useState([]);

  useEffect(() => {
    // Fetch projects from the backend API
    fetch('/projects')
      .then((res) => res.json())
      .then((data) => setProjects(data));
  }, []);

  return (
    <div>
      <Navbar />
      <main className="p-4">
        <h1 className="text-2xl font-bold">My Projects</h1>
        <ProjectGrid projects={projects} />
      </main>
    </div>
  );
}

export default Projects;
import React from 'react';

function ProjectGrid({ projects }) {
  return (
    <div className="grid grid-cols-1 md:grid-cols-2 gap-4 p-4">
      {projects.map((project) => (
        <div key={project.id} className="border p-4 rounded-lg shadow-md">
          <h2 className="text-lg font-bold">{project.title}</h2>
          <p>{project.description}</p>
        </div>
      ))}
    </div>
  );
}

export default ProjectGrid;
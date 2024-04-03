import React from 'react';

type Prop = {
  title: string;
  isActive?: boolean;
}
export const Head =({ title, isActive = true }: Prop) => {
  return (
    <div>
      <h1>Hello</h1>
      {isActive && <h3>{title}</h3>}
    </div>
  );
}
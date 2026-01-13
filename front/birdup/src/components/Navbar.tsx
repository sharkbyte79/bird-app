// import React, { useState, useEffect } from "react";
import React from "react";
import { Link } from "react-router-dom";

export interface NavbarProps {}

export default function Navbar({}: NavbarProps) {
  // const [username, setUsername] = useState<string>("Guest");

  return (
    <header className="sticky bg-white flex top-0 z-40 w-full">
      {/* container around all header bar components */}
      <div className="flex mx-auto gap-8 max-w-6xl px-2 sm:px-6 lg:px-8">
        {/* left side of header with clickable logo */}
        <span>
          <a href="/" className="text-blue-400 text-3xl">
            BirdUp
          </a>
        </span>
        {/* center of header with navigational links */}
        <nav>
          <ul className="flex flex-row gap-12">
            <Link to="/">Search Observations</Link>
            <Link to="/">Following</Link>
          </ul>
        </nav>
      </div>
    </header>
  );
}

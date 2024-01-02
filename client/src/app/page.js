"use client";

import React from 'react';
import ChatComponent from './components/chat';

const Home = () => {
  return (
    <div className='flex flex-col justify-center items-center w-screen h-screen'>
      <h1>Asisten Virtual</h1>
      <div className='relative w-[95%] h-[95%] border-solid -mt-2 px-[2%] py-[2%]'>
        <ChatComponent />
      </div>
    </div>
  );
};

export default Home;

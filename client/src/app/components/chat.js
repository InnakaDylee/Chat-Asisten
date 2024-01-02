import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Input } from './input';

const ChatComponent = () => {
  const [messages, setMessages] = useState([]);
  const [inputMessage, setInputMessage] = useState('');

  useEffect(() => {
    fetchConversationHistory();
  }, []);

  const fetchConversationHistory = async () => {
    try {
      const response = await axios.get('http://localhost:8080/chat-asisten');
      console.log(response.data); // Log respons untuk memeriksa struktur datanya

      // Pastikan bahwa 'data' ada di dalam respons dan merupakan array
      if (response.data && Array.isArray(response.data.data)) {
        setMessages(response.data.data); // Menyimpan 'data' dari respons sebagai messages
      } else {
        console.error('Invalid data format in API response');
      }
    } catch (error) {
      console.error('Error fetching conversation history:', error);
    }
  };
  
  
  const sendMessage = async () => {
    if (inputMessage.trim() !== '') {
      const newMessage = {
        role: 'user',
        content: inputMessage,
      };
      setMessages([...messages, newMessage]);
      setInputMessage('');

      try {
        const response = await axios.post('http://localhost:8080/chat-asisten', {
            role: 'user',
            content: inputMessage,
        });
        setMessages([...messages, response.data.data]);
      } catch (error) {
        console.error('Error sending message:', error);
      }
    }
  };

  return (
    <div className="h-full border-solid gap-4 flex flex-col">
      <div className="flex-1 h-[90%] border-solid rounded-[45px] overflow-y-auto p-4 bg-slate-700 bg-opacity-60 ">
        {messages.map((msg, index) => (
          <div
            key={index}
            className={`message ${msg.role === 'user' ? 'text-right' : 'text-left'} p-2`}
          >
            <span className={`message ${msg.role === 'user' ? 'float-right text-slate-300' : 'float-left text-slate-300'}`}>
                {msg.role}
            </span>
            <br/>
            <span className={` ${msg.role === 'user' ? 'bg-sky-900' : 'bg-sky-950'} text-justify inline-block break-words mx-auto min-w-16 max-w-[50%] p-2 rounded-lg my-2`}>
                {msg.content}
            </span>
          </div>
        ))}
      </div>
      <div className="flex flex-row h-[10%] justify-center items-center gap-6 p-4">
        <Input value={inputMessage} onChange={(e) => setInputMessage(e.target.value)} />
        <button className=' w-max h-max p-3 bg-gray-600 bg-opacity-40 rounded-[30px] text-gray-400' onClick={sendMessage}>Kirim</button>
      </div>
    </div>
  );
};

export default ChatComponent;
    
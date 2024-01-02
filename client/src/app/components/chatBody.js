import React from 'react';

const ChatBody = ({ data }) => {
  return (
    <div className="h-full border-solid gap-6">
      {data.map((message, index) => (
        <div key={index} className={message.sender === 'user' ? 'text-right' : 'text-left'}>
          <div className="text-sm">{message.sender}</div>
          <div className={message.sender === 'user' ? 'bg-blue text-white' : 'bg-grey text-dark-secondary'}>
            {message.content}
          </div>
        </div>
      ))}
    </div>
  );
};

export default ChatBody;

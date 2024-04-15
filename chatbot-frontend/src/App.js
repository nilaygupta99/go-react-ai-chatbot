import React, { useState, useRef, useEffect } from 'react';
import './App.css';

function App() {
  const [userInput, setUserInput] = useState('');
  const [selectedImage, setSelectedImage] = useState(null);
  const [chatHistory, setChatHistory] = useState([]);
  const [isLoading, setIsLoading] = useState(false);

  const chatEndRef = useRef(null);

  const scrollToBottom = () => {
    chatEndRef.current.scrollIntoView({ behavior: 'smooth' });
  };

  useEffect(() => {
    scrollToBottom();
  }, [chatHistory]);

  const sendMessage = () => {
    if (!userInput.trim() && !selectedImage) return;

    setChatHistory(prevChatHistory => [...prevChatHistory, { sender: 'user', message: userInput }]);
    setUserInput('');
    setIsLoading(true);

    if (selectedImage) {
      uploadImage(selectedImage, userInput)
        .then(imageData => {
          setChatHistory(prevChatHistory => [...prevChatHistory, { sender: 'user', imageUrl: URL.createObjectURL(selectedImage) }]);
          setChatHistory(prevChatHistory => [...prevChatHistory, { sender: 'bot', messages: imageData.responses }]);
          setIsLoading(false);
        })
        .catch(error => {
          console.error('Error uploading image:', error);
          setIsLoading(false);
        });
    } else {
      sendMessageToBackend(userInput);
    }
  };

  const sendMessageToBackend = (message) => {
    fetch('http://localhost:8080/chat', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ message })
    })
    .then(response => response.json())
    .then(data => {
      setChatHistory(prevChatHistory => [...prevChatHistory, { sender: 'bot', messages: data.responses }]);
      setIsLoading(false);
    })
    .catch(error => {
      console.error('Error:', error);
      setIsLoading(false);
    });
  };

  const uploadImage = (imageFile, textData) => {
    const formData = new FormData();
    formData.append('image', imageFile);
    formData.append('text', textData);

    return fetch('http://localhost:8080/upload', {
      method: 'POST',
      body: formData
    })
    .then(response => {
      if (!response.ok) {
        throw new Error('Failed to upload image');
      }
      return response.json();
    });
  };

  const handleImageChange = (e) => {
    setSelectedImage(e.target.files[0]);
  };

  const handleKeyPress = (e) => {
    if (e.key === 'Enter') {
      sendMessage();
    }
  };

  const searchImage = () => {
    setIsLoading(true);
    fetch('http://localhost:8080/imagesearch', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ message: userInput })
    })
    .then(response => response.json())
    .then(data => {
      setChatHistory(prevChatHistory => [...prevChatHistory, { sender: 'user', message: userInput }]);
      setChatHistory(prevChatHistory => [...prevChatHistory, { sender: 'bot', imageUrl: data.image }]);
      setIsLoading(false);
    })
    .catch(error => {
      console.error('Error searching image:', error);
      setIsLoading(false);
    });
  };

  return (
    <div className="container">
      <div className="header">
        <h1>React ChatBot</h1>
      </div>
      <div className="chat-container">
        {chatHistory.map((chat, index) => (
          <div key={index} className={`message ${chat.sender}`}>
            {chat.message && <div className="content">{chat.message}</div>}
            {chat.imageUrl && <img src={chat.imageUrl} alt="Uploaded" className="content" />}
            {chat.messages && chat.messages.map((message, idx) => (
              <div key={idx} className="content">{message}</div>
            ))}
          </div>
        ))}
        <div ref={chatEndRef} />
      </div>
      <div className="input-container">
        <input
          type="text"
          value={userInput}
          onChange={(e) => setUserInput(e.target.value)}
          onKeyDown={handleKeyPress}
          placeholder="Type your message here..."
        />
        <input
          type="file"
          accept="image/*"
          onChange={handleImageChange}
        />
        <button onClick={sendMessage} disabled={isLoading}>
          {isLoading ? (
            <>
              <span>Sending...</span>
              <div className="spinner" />
            </>
          ) : (
            <span>Send</span>
          )}
        </button>
        <button onClick={searchImage} disabled={isLoading}>
          {isLoading ? (
            <>
              <span>Searching...</span>
              <div className="spinner" />
            </>
          ) : (
            <span>Image Generation</span>
          )}
        </button>
      </div>
    </div>
  );
}

export default App;

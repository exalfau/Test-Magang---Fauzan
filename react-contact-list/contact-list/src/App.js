import React, { useState } from 'react';
import './App.css';

const initialContacts = [
  { id: 1, name: 'John Doe', phone: '1234567890', email: 'john@example.com' },
  { id: 2, name: 'Jane Smith', phone: '0987654321', email: 'jane@example.com' },
];

const App = () => {
  const [contacts, setContacts] = useState(initialContacts);
  const [name, setName] = useState('');
  const [phone, setPhone] = useState('');
  const [email, setEmail] = useState('');
  const [editedContact, setEditedContact] = useState(null); // Tambahkan variabel editedContact

  const handleAddContact = () => {
    const newContact = {
      id: contacts.length + 1,
      name,
      phone,
      email,
    };
    setContacts([...contacts, newContact]);
    setName('');
    setPhone('');
    setEmail('');
  };

  const handleEditContact = (id) => {
    const editedContact = contacts.find((contact) => contact.id === id);
    if (editedContact) {
      setEditedContact(editedContact); // Update editedContact
      setName(editedContact.name);
      setPhone(editedContact.phone);
      setEmail(editedContact.email);
    }
  };

  const handleUpdateContact = () => {
    const updatedContacts = contacts.map((contact) => {
      if (contact.id === editedContact.id) {
        return { ...contact, name, phone, email };
      }
      return contact;
    });
    setContacts(updatedContacts);
    setName('');
    setPhone('');
    setEmail('');
    setEditedContact(null); // Reset editedContact setelah update
  };

  const handleDeleteContact = (id) => {
    const updatedContacts = contacts.filter((contact) => contact.id !== id);
    setContacts(updatedContacts);
  };

  return (
    <div className="app">
      <h1>Daftar Kontak</h1>
      <div className="contact-form">
        <input
          type="text"
          placeholder="Nama"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
        <input
          type="text"
          placeholder="Nomor Telepon"
          value={phone}
          onChange={(e) => setPhone(e.target.value)}
        />
        <input
          type="text"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <button onClick={handleAddContact}>Tambah Kontak</button>
        {editedContact && ( // Tampilkan tombol Update hanya jika ada editedContact
          <button onClick={handleUpdateContact}>Update Kontak</button>
        )}
      </div>
      <div className="contact-list">
        {contacts.map((contact) => (
          <div key={contact.id} className="contact-item">
            <div>
              <strong>{contact.name}</strong>
            </div>
            <div>{contact.phone}</div>
            <div>{contact.email}</div>
            <div>
              <button onClick={() => handleEditContact(contact.id)}>Edit</button>
              <button onClick={() => handleDeleteContact(contact.id)}>Hapus</button>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default App;

import React, { useState } from 'react';

const productsData = [
  {
    id: 1,
    name: 'Xiaomi 13 Ultra',
    description: 'The device is protected with extra seals to prevent failures caused by dust, raindrops, and water splashes.',
    price: 800,
    image: 'https://images.versus.io/objects/xiaomi-13-ultra.front.medium.1681827673875.webp',
  },
  {
    id: 2,
    name: 'Oppo Find X5 Pro 5G',
    description: 'A newer version of Android generally provides a better user experience. Major Android releases include new features and improvements to performance, as well as important security updates.',
    price: 999,
    image: 'https://images.versus.io/objects/oppo-find-x5-pro-5g.front.medium.1644238019281.webp',
  },   
  {
    id: 3,
    name: 'Asus Zenfone 9',
    description: 'Optical image stabilization uses gyroscopic sensors to detect the vibrations of the camera. The lens adjusts the optical path accordingly, ensuring that any type of motion blur is corrected before the sensor captures the image.',
    price: 750,
    image: 'https://images.versus.io/objects/asus-zenfone-9.front.medium.1659038869878.webp',
  },
];

const ProductList = () => {
  const [products, setProducts] = useState(productsData);
  const [sortOption, setSortOption] = useState('');

  const handleSortChange = (event) => {
    const value = event.target.value;
    setSortOption(value);

    let sortedProducts = [...products];

    if (value === 'ascending-price') {
      sortedProducts.sort((a, b) => a.price - b.price);
    } else if (value === 'descending-price') {
      sortedProducts.sort((a, b) => b.price - a.price);
    } else if (value === 'ascending-name') {
      sortedProducts.sort((a, b) => a.name.localeCompare(b.name));
    } else if (value === 'descending-name') {
      sortedProducts.sort((a, b) => b.name.localeCompare(a.name));
    }

    setProducts(sortedProducts);
  };

  return (
    <div>
      <label>
        Sort By:
        <select value={sortOption} onChange={handleSortChange}>
          <option value="">None</option>
          <option value="ascending-price">Price (Low to High)</option>
          <option value="descending-price">Price (High to Low)</option>
          <option value="ascending-name">Name (A to Z)</option>
          <option value="descending-name">Name (Z to A)</option>
        </select>
      </label>
      <ul>
        {products.map((product) => (
          <li key={product.id}>
            <img src={product.image} alt={product.name} />
            <h3>{product.name}</h3>
            <p>{product.description}</p>
            <p>Price: ${product.price}</p>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default ProductList;

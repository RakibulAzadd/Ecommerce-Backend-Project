import React from 'react';
import ProductCard from './ProductCard';
import '../styles/ProductList.css';

// function ProductList({ products, onEdit, onDelete }) {
//   return (
//     <div className="product-list">
//       {products.map((product) => (
//         <ProductCard
//           key={product._id}
//           product={product}
//           onEdit={onEdit}
//           onDelete={onDelete}
//         />
//       ))}
//     </div>
//   );
// }


// ProductList.js
function ProductList({ products, onEdit, onDelete }) {
  console.log(products)
  return (
    <div className="product-list">
      {products.map((product) => (
        <ProductCard
          key={product.id} // Use product.id if that's what your backend returns
          product={product}
          onEdit={() => onEdit(product)}
          onDelete={() => onDelete(product.id)} // Pass the id to onDelete
        />
      ))}
    </div>
  );
}

export default ProductList;
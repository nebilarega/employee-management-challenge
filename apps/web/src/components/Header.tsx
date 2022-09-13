import React from 'react';
interface Adding {
  setIsAdding(isadding: boolean): void;
}

const Header: React.FC<Adding> = ({ setIsAdding }) => {
  return (
    <header>
      <h1 className="text-center">Employee Management Software</h1>
      <div className="mt-[30px] mb-[18px] flex justify-end">
        <button onClick={() => setIsAdding(true)} className="rounded-full ">
          Add Button
        </button>
      </div>
    </header>
  );
};

export default Header;

import React from "react";
import Navbar from "../component/navbar";
import Headers from "../component/header";
import Footer from "../component/footer";
import Menu from "../component/menu";

const Subject = () => {
  return (
    <>
      <Headers />
      <Navbar />
      <div className="container text-center">
        <div className="row">
          <p>รายวิชาทั้งหมด</p>
        </div>
      </div>
      <Footer/>
    </>
  );
};

export default Subject;

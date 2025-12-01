import React from "react";
import Headers from "../component_admin/header";
import Navbar from "../component_admin/navbar";
import Menu from "../component_admin/menu";
import Footer from "../component_admin/footer";
import { Outlet } from "react-router-dom";

const AdminLayout = () => {
  return (
    <>
        <Headers />
        <Navbar />
      <div className="container text-center">
        <div className="row">
          <div className="col-sm-3">
            <Menu />
          </div>
          <div className="col-sm-9 text-start">

          </div>
        </div>
      </div>
          <Footer />
    </>
  );
};

export default AdminLayout;

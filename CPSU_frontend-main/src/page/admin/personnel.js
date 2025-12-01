import React, { useEffect, useState } from "react";
import "../../css/admin/personnel.css";
import { Link } from "react-router-dom";
import axios from "axios";
import Headers from "../../component/header";
import Navbar from "../../component/navbar";
import Footer from "../../component/footer";
import Menu from "../../component/menu";

function Personnel_admin() {
  const [personnel, setPersonnel] = useState([]);

  useEffect(() => {
    const fetchPersonnel = async () => {
      try {
        const res = await axios.get(
          "http://localhost:8080/api/v1/admin/personnel"
        );
        setPersonnel(res.data);
      } catch (err) {
        console.error("Error fetching personnel", err);
      }
    };

    fetchPersonnel();
  }, []);

  return (
    <>
      <Headers />
      <Navbar />
      <div className="container text-center mt-4">
        <div className="row">
          <div className="col-sm-3">
            <Menu />
          </div>

          <div className="col-sm-9">
            <div className="d-flex justify-content-between align-items-center mb-4">
              <h3 id="personnel-title">บุคลากรทั้งหมด</h3>
              <Link to="/admin/addpersonnel" className="btn-addpersonnel">
                 + เพิ่มบุคลากร
              </Link>
            </div>

            <div className="row row-cols-1 row-cols-md-3 g-4">
              {personnel.map((item, index) => (
                <div className="col" key={index}>
                  <Link
                    to={`/admin/personnel/${item.personnel_id}`}
                    className="text-decoration-none text-dark"
                  >
                    <div className="card h-100 shadow-sm">
                      <img
                        src={item.file_image}
                        className="card-img-top personnel-img"
                        alt={item.thai_name}
                        onError={(e) =>
                          (e.target.src = "/images/default-profile.png")
                        }
                      />
                      <div className="card-body">
                        <h5 className="card-title name-title">
                          {item.thai_academic_position} {item.thai_name}
                        </h5>
                      </div>
                    </div>
                  </Link>
                </div>
              ))}
            </div>
          </div>
        </div>
      </div>

      <div id="space"></div>
      <Footer />
    </>
  );
}

export default Personnel_admin;

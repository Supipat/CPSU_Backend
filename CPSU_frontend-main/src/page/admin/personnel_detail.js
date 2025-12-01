import React, { useEffect, useState } from "react";
import { useParams, Link } from "react-router-dom";
import axios from "axios";
import "../../css/admin/news.css";
import Headers from "../../component/header";
import Navbar from "../../component/navbar";
import Footer from "../../component/footer";
import Menu from "../../component/menu";

function Detail_Personnel() {
  const { id } = useParams();
  const [person, setPerson] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchPerson = async () => {
      try {
        const res = await axios.get(
          `http://localhost:8080/api/v1/admin/personnel/${id}`
        );
        setPerson(res.data.personnel);
        setLoading(false);
      } catch (err) {
        console.error("Error fetching personnel detail:", err);
        setError(err);
        setLoading(false);
      }
    };

    fetchPerson();
  }, [id]);

  if (loading) {
    return (
      <>
        <Headers />
        <Navbar />
        <div className="container text-center mt-5">
          <p>กำลังโหลดข้อมูล...</p>
        </div>
        <Footer />
      </>
    );
  }

  if (error || !person) {
    return (
      <>
        <Headers />
        <Navbar />
        <div className="container text-center mt-5">
          <p className="text-danger">ไม่สามารถโหลดข้อมูลบุคลากรได้</p>
        </div>
        <Footer />
      </>
    );
  }

  return (
    <>
      <Headers />
      <Navbar />
      <div className="container-fluid">
        <div className="row">
          <div className="col-sm-3">
            <Menu />
          </div>
          <div className="col-sm-9">
            {" "}
            <div className="d-grid gap-2 d-md-flex justify-content-md-end">
              <div className="edit_detele-position">
                {person && (
                  <Link
                    to={`/admin/editpersonnel/${id}`}
                    className="btn btn-warning"
                    id="btn-edit"
                  >
                    แก้ไข
                  </Link>
                )}
                <div className="btn btn-danger" id="btn-delete">
                  ลบ
                </div>
              </div>
            </div>
            <div className="card mt-4 shadow-sm p-4">
              <div className="row">
                <div className="col-md-4 text-center">
                  <img
                    src={person.file_image}
                    alt={person.thai_name}
                    className="img-fluid rounded border"
                    onError={(e) =>
                      (e.target.src = "/images/default-profile.png")
                    }
                  />
                  <p className="mt-3">
                    <strong>อีเมล:</strong> <br />
                    <a href={`mailto:${person.email}`}>{person.email}</a>
                  </p>
                  <p>
                    <strong>เว็บไซต์:</strong> <br />
                    <a href={person.website} target="_blank" rel="noreferrer">
                      {person.website}
                    </a>
                  </p>
                </div>
                <div className="col-md-8">
                  <h4>
                    {person.thai_academic_position} {person.thai_name}
                  </h4>
                  <p className="text-muted">
                    {person.eng_academic_position} {person.eng_name}
                  </p>
                  <p>
                    <strong>ตำแหน่งในภาควิชา:</strong>{" "}
                    {person.department_position_name}
                  </p>
                  <p>
                    <strong>ประเภทบุคลากร:</strong> {person.type_personnel}
                  </p>

                  <hr />

                  <h5>ประวัติการศึกษา</h5>
                  <ul>
                    {person.education?.split("\n").map((line, i) => (
                      <li key={i}>{line.trim()}</li>
                    ))}
                  </ul>

                  <h5 className="mt-3">สาขาที่เชี่ยวชาญ</h5>
                  <ul>
                    {person.related_fields?.split("\n").map((line, i) => (
                      <li key={i}>{line.trim()}</li>
                    ))}
                  </ul>
                </div>
              </div>

              <div className="text-end mt-4">
                <Link
                  to="/admin/personnel"
                  className="btn btn-secondary btn-sm"
                >
                  ← กลับไปหน้ารายการบุคลากร
                </Link>
              </div>
            </div>
          </div>
        </div>
      </div>
      <Footer />
    </>
  );
}

export default Detail_Personnel;

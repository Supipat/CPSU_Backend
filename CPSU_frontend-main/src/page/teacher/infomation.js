import React, { useState } from "react";
import Footer from "../../component/footer";
import Header from "../../component/header";
import Navbar from "../../component/navbar";
import { useNavigate } from "react-router-dom";

const TeacherInfomation = () => {
  const [copied, setCopied] = useState(false);
  const navigate = useNavigate();

  // ตัวอย่าง handler เมื่อกดปุ่มแก้ไข
  const handleEdit = () => {
    navigate("/teacher/informationedit"); // <-- ใช้ navigate เปลี่ยนหน้า
  };

  const handleCopyEmail = () => {
    navigator.clipboard.writeText(data.email).then(() => {
      setCopied(true);
      setTimeout(() => setCopied(false), 2000);
    });
  };

  const data = {
    personnel_id: 1,
    type_personnel: "สายวิชาการ",
    department_position_id: 1,
    department_position_name: "หัวหน้าภาควิชา",
    academic_position_id: 2,
    thai_academic_position: "ผศ.ดร.",
    eng_academic_position: "Asst.Prof.Dr.",
    thai_name: "สิรักข์ แก้วจำนงค์",
    eng_name: "Sirak Kaewjamnong",
    education: `Ph.D. (Computer Science) Lancaster University, UK (2015)
วศ.ม. (วิศวกรรมคอมพิวเตอร์) มหาวิทยาลัยเกษตรศาสตร์ (2544)
วท.บ. (วิทยาการคอมพิวเตอร์) สถาบันเทคโนโลยีราชมงคล (2540)`,
    related_fields: `Computer Network Architectures
Algorithms and Protocols`,
    email: "kaewjamnong_s@su.ac.th",
    website: "https://webserv.cp.su.ac.th/lecturer/sirak",
    file_image:
      "https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/personnel/Sirak.jpg",
  };

  return (
    <>
      <Header />
      <Navbar />
      <h4 className="text-center my-4">บุคลากร</h4>
      <div className="container my-5">
        <div className="teacher-card mx-auto bg-white shadow rounded p-4 position-relative">
          {/* ปุ่มแก้ไขตำแหน่งมุมบนขวา */}
          <button
            onClick={handleEdit}
            className="btn btn-outline-secondary position-absolute top-0 end-0 m-3 d-flex align-items-center gap-1"
            title="แก้ไขข้อมูล"
            style={{ zIndex: 10 }}
          >
            <i className="bi bi-pencil"></i> แก้ไข
          </button>

          <div className="row">
            {/* Left: Image and Email */}
            <div className="col-md-4 text-center mb-4 mb-md-0">
              <img
                src={data.file_image}
                alt={data.thai_name}
                className="img-fluid img-profile"
                onError={(e) => (e.target.src = "/images/default-profile.png")}
              />

              <div className="d-flex justify-content-center align-items-center gap-2 mt-3 flex-wrap">
                <div>
                  <strong>อีเมล</strong>
                  <br />
                  <a href={`mailto:${data.email}`}>{data.email}</a>
                </div>
                <button
                  className="btn btn-outline-primary btn-sm d-flex align-items-center gap-1 copy-btn"
                  onClick={handleCopyEmail}
                  title="คัดลอกอีเมล"
                >
                  <i className="bi bi-clipboard"></i> คัดลอก
                </button>
              </div>
              {copied && (
                <small className="text-success d-block mt-1">คัดลอกแล้ว!</small>
              )}
            </div>

            {/* Right: Info */}
            <div className="col-md-8 text-start">
              <h5 className="text-primary">{data.department_position_name}</h5>
              <hr />
              <h4 className="fw-bold mt-2">
                {data.thai_academic_position} {data.thai_name}
              </h4>
              <p className="text-muted mb-1">
                {data.eng_academic_position} {data.eng_name}
              </p>
              <hr />
              <h5 className="mt-4">ประวัติการศึกษา</h5>
              <ul>
                {data.education
                  .split("\n")
                  .filter((line) => line.trim() !== "")
                  .map((line, index) => (
                    <li key={index}>{line.trim()}</li>
                  ))}
              </ul>
              <hr />
              <h5 className="mt-4">สาขาที่เชี่ยวชาญ</h5>
              <ul>
                {data.related_fields
                  .split("\n")
                  .filter((line) => line.trim() !== "")
                  .map((line, index) => (
                    <li key={index}>{line.trim()}</li>
                  ))}
              </ul>
              <hr />
              <p className="mt-3">
                <strong>เว็บไซต์ </strong>{" "}
                <a href={data.website} target="_blank" rel="noreferrer">
                  {data.website}
                </a>
              </p>
            </div>
          </div>

          {/* Office Hours */}
          <hr className="my-4" />
          <h5 className="text-center mb-3">เวลาที่สะดวก (Office Hours)</h5>
          <div className="table-responsive">
            <table className="table table-bordered text-center office-hours">
              <thead className="table-light">
                <tr>
                  <th>จันทร์</th>
                  <th>อังคาร</th>
                  <th>พุธ</th>
                  <th>พฤหัสบดี</th>
                  <th>ศุกร์</th>
                  <th>เสาร์</th>
                  <th>อาทิตย์</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td>10:00 - 12:00</td>
                  <td>13:00 - 15:00</td>
                  <td>10:00 - 12:00</td>
                  <td>13:00 - 15:00</td>
                  <td>09:00 - 11:00</td>
                  <td>-</td>
                  <td>-</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>

      <div className="teacher-research mx-auto bg-white shadow-sm rounded p-4 mt-4">
        <h5 className="mb-3">ผลงานวิจัย (Research Publications)</h5>
        <ul className="research-list">
          <li>
            Suphachai Phawiakkharakun, Sunee Pongpinigpinyo, (2024).{" "}
            <em>
              Enhanced non-destructive of degree of pineapple juiciness using
              ensemble learning model based on tapping sound sensing
            </em>
            , International Journal of Applied Science and Engineering, Volume
            21, No. 1, March 2024.
          </li>
          <li className="mt-3">
            Zhang, Y., Li, X., Wang, J., & Chen, Q. (2023).{" "}
            <em>
              Deep learning-based method for early diagnosis of lung cancer
              using CT images
            </em>
            . Computers in Biology and Medicine, 153, 106388.{" "}
          </li>
          <li className="mt-3">
            Smith, A., Johnson, R., & Lee, M. (2022).{" "}
            <em>
              Optimization of photovoltaic panels using hybrid metaheuristic
              algorithms
            </em>
            . Renewable Energy, 185, 757-769.{" "}
          </li>
          <li className="mt-3">
            Garcia, L., Kumar, P., & Martinez, F. (2021).{" "}
            <em>
              A novel IoT-based smart irrigation system for precision
              agriculture
            </em>
            . Sensors, 21(18), 6200.{" "}
          </li>
        </ul>
      </div>

      <Footer />
    </>
  );
};

export default TeacherInfomation;

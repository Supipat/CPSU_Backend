import React, { useState } from "react";
import Footer from "../../component/footer";
import Header from "../../component/header";
import Navbar from "../../component/navbar";
import { useNavigate } from "react-router-dom";

const EditTeacherInformation = () => {
  // ข้อมูล mock ของครู
  const initialData = {
    personnel_id: 1,
    type_personnel: "สายวิชาการ",
    department_position_name: "หัวหน้าภาควิชา",
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

  const [formData, setFormData] = useState(initialData);
  const navigate = useNavigate();

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    alert("บันทึกข้อมูลเรียบร้อย");
    navigate("/teacher/information");
  };

  const handleCancel = () => {
    // reset ข้อมูลกลับเป็นต้นฉบับ
    setFormData(initialData);
  };

  return (
    <>
      <Header />
      <Navbar />
      <div className="container my-5">
        <h3 className="mb-4">แก้ไขข้อมูลบุคลากร</h3>
        <form onSubmit={handleSubmit}>
          <div className="mb-3">
            <label className="form-label">ตำแหน่งภาควิชา</label>
            <input
              type="text"
              className="form-control"
              name="department_position_name"
              value={formData.department_position_name}
              onChange={handleChange}
              required
            />
          </div>

          <div className="mb-3">
            <label className="form-label">ตำแหน่งวิชาการ (ไทย)</label>
            <input
              type="text"
              className="form-control"
              name="thai_academic_position"
              value={formData.thai_academic_position}
              onChange={handleChange}
            />
          </div>

          <div className="mb-3">
            <label className="form-label">ตำแหน่งวิชาการ (อังกฤษ)</label>
            <input
              type="text"
              className="form-control"
              name="eng_academic_position"
              value={formData.eng_academic_position}
              onChange={handleChange}
            />
          </div>

          <div className="mb-3">
            <label className="form-label">ชื่อ (ไทย)</label>
            <input
              type="text"
              className="form-control"
              name="thai_name"
              value={formData.thai_name}
              onChange={handleChange}
              required
            />
          </div>

          <div className="mb-3">
            <label className="form-label">ชื่อ (อังกฤษ)</label>
            <input
              type="text"
              className="form-control"
              name="eng_name"
              value={formData.eng_name}
              onChange={handleChange}
            />
          </div>

          <div className="mb-3">
            <label className="form-label">ประวัติการศึกษา</label>
            <textarea
              className="form-control"
              rows={4}
              name="education"
              value={formData.education}
              onChange={handleChange}
            ></textarea>
            <small className="form-text text-muted">
              กรอกข้อมูลโดยเว้นบรรทัดใหม่สำหรับแต่ละบรรทัด
            </small>
          </div>

          <div className="mb-3">
            <label className="form-label">สาขาที่เชี่ยวชาญ</label>
            <textarea
              className="form-control"
              rows={2}
              name="related_fields"
              value={formData.related_fields}
              onChange={handleChange}
            ></textarea>
            <small className="form-text text-muted">
              กรอกข้อมูลโดยเว้นบรรทัดใหม่สำหรับแต่ละบรรทัด
            </small>
          </div>

          <div className="mb-3">
            <label className="form-label">อีเมล</label>
            <input
              type="email"
              className="form-control"
              name="email"
              value={formData.email}
              onChange={handleChange}
              required
            />
          </div>

          <div className="mb-3">
            <label className="form-label">เว็บไซต์</label>
            <input
              type="url"
              className="form-control"
              name="website"
              value={formData.website}
              onChange={handleChange}
            />
          </div>

          <div className="mb-3">
            <label className="form-label">URL รูปภาพ</label>
            <input
              type="url"
              className="form-control"
              name="file_image"
              value={formData.file_image}
              onChange={handleChange}
            />
          </div>

          <button type="submit" className="btn btn-primary me-2">
            บันทึก
          </button>
          <button
            type="button"
            onClick={handleCancel}
            className="btn btn-secondary"
          >
            ยกเลิก
          </button>
        </form>
      </div>
      <Footer />
    </>
  );
};

export default EditTeacherInformation;

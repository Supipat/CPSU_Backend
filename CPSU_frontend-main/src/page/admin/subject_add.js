import React, { useState, useEffect } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import Headers from "../../component/header";
import Navbar from "../../component/navbar";
import Footer from "../../component/footer";
import Menu from "../../component/menu";
import "../../css/admin/subject.css";

const Subject_add = () => {
  const navigate = useNavigate();
  const [courseOptions, setCourseOptions] = useState([]);

  useEffect(() => {
    const fetchCourses = async () => {
      try {
        const response = await axios.get(
          "http://localhost:8080/api/v1/admin/course"
        );
        setCourseOptions(response.data);
      } catch (error) {
        console.error("Error loading courses", error);
      }
    };

    fetchCourses();
  }, []);

  const [formData, setFormData] = useState({
    subject_id: "",
    thai_subject: "",
    eng_subject: "",
    credits: "",
    semester_id: "",
    compulsory_subject: "",
    condition: "",
    description_thai: "",
    description_eng: "",
    clo: "",
    course_id: "",
    plan_type_id: "",
  });

  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");
    setSuccess("");

    try {
      const cleanData = {
        subject_id: formData.subject_id,
        thai_subject: formData.thai_subject,
        eng_subject: formData.eng_subject,
        credits: formData.credits,
        course_id: formData.course_id,
        semester_id: parseInt(formData.semester_id),
        compulsory_subject:
          formData.compulsory_subject.trim() === ""
            ? "-"
            : formData.compulsory_subject,
        condition: formData.condition.trim() === "" ? "-" : formData.condition,
        description_thai: formData.description_thai,
        description_eng: formData.description_eng,
        clo: formData.clo.trim() === "" ? "-" : formData.clo,
        plan_type_id: parseInt(formData.plan_type_id),
      };

      console.log(cleanData);
      await axios.post("http://localhost:8080/api/v1/admin/subject", cleanData);
      setSuccess("เพิ่มรายวิชาสำเร็จ!");
      setFormData({
        subject_id: "",
        thai_subject: "",
        eng_subject: "",
        credits: "",
        semester_id: "",
        compulsory_subject: "",
        condition: "",
        description_thai: "",
        description_eng: "",
        clo: "",
        course_id: "",
        plan_type_id: "",
      });

      navigate("/admin/subject");
    } catch (err) {
      console.error(err);
      setError("เกิดข้อผิดพลาดในการเพิ่มรายวิชา");
    }
  };

  return (
    <>
      <Headers />
      <Navbar />
      <div className="container mt-4">
        <div className="row">
          <div className="col-sm-3">
            <Menu />
          </div>
          <div className="col-sm-9 text-start">
            <h2>เพิ่มรายวิชาใหม่</h2>
            <form onSubmit={handleSubmit}>
              <div className="mb-3">
                <label>หลักสูตร (Course)</label>
                <select
                  name="course_id"
                  value={formData.course_id}
                  onChange={handleChange}
                  className="form-control"
                  required
                >
                  <option value="">-- เลือกหลักสูตร --</option>
                  {courseOptions.map((course) => (
                    <option key={course.course_id} value={course.course_id}>
                      {course.thai_course}
                    </option>
                  ))}
                </select>
              </div>
              <div className="mb-3">
                <label>รหัสรายวิชา</label>
                <input
                  type="text"
                  name="subject_id"
                  value={formData.subject_id}
                  onChange={handleChange}
                  className="form-control"
                  required
                  placeholder="เช่น XXXXXX , SUXXX"
                />
              </div>

              <div className="mb-3">
                <label>ชื่อรายวิชา (ภาษาไทย)</label>
                <input
                  type="text"
                  name="thai_subject"
                  value={formData.thai_subject}
                  onChange={handleChange}
                  className="form-control"
                  required
                />
              </div>

              <div className="mb-3">
                <label>ชื่อรายวิชา (ภาษาอังกฤษ)</label>
                <input
                  type="text"
                  name="eng_subject"
                  value={formData.eng_subject}
                  onChange={handleChange}
                  className="form-control"
                />
              </div>

              <div className="mb-3">
                <label>หน่วยกิต</label>
                <input
                  type="text"
                  name="credits"
                  value={formData.credits}
                  onChange={handleChange}
                  className="form-control"
                  required
                  placeholder="เช่น X , X(X-X-X)"
                />
              </div>

              <div className="mb-3">
                <label>ชั้นปีและภาคการศึกษา</label>
                <select
                  name="semester_id"
                  value={formData.semester_id}
                  onChange={handleChange}
                  className="form-control"
                  required
                >
                  <option value="">-- เลือกภาคการศึกษา --</option>
                  <option value="11">ปีที่ 1 ภาคการศึกษาที่ 1</option>
                  <option value="12">ปีที่ 1 ภาคการศึกษาที่ 2</option>
                  <option value="21">ปีที่ 2 ภาคการศึกษาที่ 1</option>
                  <option value="22">ปีที่ 2 ภาคการศึกษาที่ 2</option>
                  <option value="31">ปีที่ 3 ภาคการศึกษาที่ 1</option>
                  <option value="32">ปีที่ 3 ภาคการศึกษาที่ 2</option>
                  <option value="41">ปีที่ 4 ภาคการศึกษาที่ 1</option>
                  <option value="42">ปีที่ 4 ภาคการศึกษาที่ 2</option>
                </select>
              </div>

              <div className="mb-3">
                <label>แผนการศึกษา</label>
                <select
                  name="plan_type_id"
                  value={formData.plan_type_id}
                  onChange={handleChange}
                  className="form-control"
                  required
                >
                  <option value="">-- เลือกแผนการศึกษา --</option>
                  <option value="1">โครงงานวิจัย</option>
                  <option value="2">สหกิจศึกษา</option>
                </select>
              </div>
              <div className="mb-3">
                <label>วิชาบังคับ (ถ้ามี)</label>
                <input
                  type="text"
                  name="compulsory_subject"
                  value={formData.compulsory_subject}
                  onChange={handleChange}
                  className="form-control"
                  placeholder="เช่น CS517XXX ชื่อรายวิชา, IT520XXX ชื่อรายวิชา"
                />
              </div>

              <div className="mb-3">
                <label>เงื่อนไข (Prerequisites)</label>
                <input
                  type="text"
                  name="condition"
                  value={formData.condition}
                  onChange={handleChange}
                  className="form-control"
                />
              </div>

              <div className="mb-3">
                <label>คำอธิบายรายวิชา (ภาษาไทย)</label>
                <textarea
                  name="description_thai"
                  value={formData.description_thai}
                  onChange={handleChange}
                  className="form-control"
                  rows={3}
                />
              </div>

              <div className="mb-3">
                <label>Course Description (English)</label>
                <textarea
                  name="description_eng"
                  value={formData.description_eng}
                  onChange={handleChange}
                  className="form-control"
                  rows={3}
                />
              </div>

              <div className="mb-3">
                <label>CLOs</label>
                <textarea
                  name="clo"
                  value={formData.clo}
                  onChange={handleChange}
                  className="form-control"
                  rows={5}
                  placeholder="CLOXX XXXXXXXXXX"
                />
              </div>

              {error && <div className="alert alert-danger">{error}</div>}
              {success && <div className="alert alert-success">{success}</div>}

              <button
                type="submit"
                className="btn btn-primary"
                id="btn-addsubject"
              >
                บันทึกรายวิชา
              </button>
            </form>
          </div>
        </div>
      </div>
      <Footer />
    </>
  );
};

export default Subject_add;

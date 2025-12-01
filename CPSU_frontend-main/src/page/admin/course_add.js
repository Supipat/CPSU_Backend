import React, { useState, useEffect } from "react";
import axios from "axios";
import Headers from "../../component/header";
import Navbar from "../../component/navbar";
import Footer from "../../component/footer";
import Menu from "../../component/menu";
import "../../css/admin/course_add.css";
import { useNavigate } from "react-router-dom";

const Add_course = () => {
  const [structureFile, setStructureFile] = useState(null);
  const [roadmapFile, setRoadmapFile] = useState(null);
  const navigate = useNavigate();

  const [formData, setFormData] = useState({
    course_id: "",
    degree_id: "",
    major_id: "",
    degree_name_id: "",
    career_paths: "",
    plo: "",
    thai_course: "",
    eng_course: "",
    thai_degree: "",
    eng_degree: "",
    year: "",
    admission_req: "",
    graduation_req: "",
    philosophy: "",
    objective: "",
    tuition: "",
    credits: "",
    detail_url: "",
    status: "แสดง",
  });

  const handleFileChange = (e) => {
    const { name, files } = e.target;
    if (name === "structure") {
      setStructureFile(files[0]);
    } else if (name === "roadmap") {
      setRoadmapFile(files[0]);
    }
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (!formData.major_id) {
      alert("กรุณาเลือกสาขาให้ถูกต้องก่อนบันทึก");
      return;
    }

    try {
      const coursePayload = {
        course_id: formData.course_id,
        degree_id: parseInt(formData.degree_id) || null,
        major_id: parseInt(formData.major_id) || null,
        degree_name_id: parseInt(formData.degree_name_id) || null,
        year: parseInt(formData.year) || 0,
        thai_course: formData.thai_course,
        eng_course: formData.eng_course,
        thai_degree: formData.thai_degree,
        eng_degree: formData.eng_degree,
        admission_req: formData.admission_req,
        graduation_req: formData.graduation_req,
        philosophy: formData.philosophy,
        objective: formData.objective,
        tuition: formData.tuition,
        credits: formData.credits,
        career_paths: formData.career_paths,
        plo: formData.plo,
        detail_url: formData.detail_url,
      };

      console.log("Payload to create course:", coursePayload);

      // ส่งข้อมูลสร้าง course

      const courseRes = await axios.post(
        "http://localhost:8080/api/v1/admin/course",
        coursePayload
      );

      if (structureFile) {
        const structureForm = new FormData();
        structureForm.append("course_structure_url", structureFile);
        structureForm.append("thai_course", structureFile.name);
        structureForm.append("course_id", courseRes.data.course_id);

        await axios.post(
          "http://localhost:8080/api/v1/admin/structure",
          structureForm,
          {
            headers: {
              "Content-Type": "multipart/form-data",
            },
          }
        );
      }

      if (roadmapFile) {
        const roadmapForm = new FormData();
        roadmapForm.append("roadmap_url", roadmapFile);
        roadmapForm.append("thai_course", roadmapFile.name);
        roadmapForm.append("course_id", courseRes.data.course_id);

        await axios.post(
          "http://localhost:8080/api/v1/admin/roadmap",
          roadmapForm,
          {
            headers: {
              "Content-Type": "multipart/form-data",
            },
          }
        );
      }

      alert("บันทึกข้อมูลหลักสูตรสำเร็จ!");
      navigate("/admin/course");
    } catch (error) {
      console.error("Error:", error.response?.data || error);
      alert("เกิดข้อผิดพลาดในการบันทึกข้อมูล");
    }
  };

  return (
    <>
      <Headers />
      <Navbar />

      <div className="container text-center">
        <div className="row">
          <div className="col-sm-3">
            <Menu />
          </div>
          <div className="col-sm-9">
            <div className="col-md-4" id="course-all">
              เพิ่มหลักสูตร
            </div>

            <form onSubmit={handleSubmit}>
              <div className="mb-3">
                <p id="text-header-coures">สถานะ (Status)</p>
                <select
                  name="status"
                  value={formData.status}
                  onChange={handleChange}
                  className="form-control"
                  required
                >
                  <option value="แสดง">แสดง</option>
                  <option value="ไม่แสดง">ไม่แสดง</option>
                </select>
              </div>

              <p id="text-header-coures">รหัสของหลักสูตร</p>
              <input
                className="form-control mb-3"
                name="course_id"
                value={formData.course_id}
                onChange={handleChange}
                required
              />

              <p id="text-header-coures">ระดับปริญญา</p>
              <select
                className="form-control mb-3"
                name="degree_id"
                value={formData.degree_id}
                onChange={(e) => {
                  setFormData((prev) => ({
                    ...prev,
                    degree_id: e.target.value,
                    degree_name_id: "",
                  }));
                }}
              >
                <option value="">-- เลือกระดับปริญญา --</option>
                <option value="1">ปริญญาตรี</option>
                <option value="2">ปริญญาโท</option>
                <option value="3">ปริญญาเอก</option>
              </select>
              <p id="text-header-coures">หลักสูตรของปี</p>
              <input
                className="form-control mb-3"
                name="year"
                value={formData.year}
                onChange={handleChange}
                type="number"
                min="1"
                required
              />

              <p id="text-header-coures">สาขา</p>
              <input
                className="form-control mb-3"
                name="major_id"
                value={formData.major_id}
                onChange={handleChange}
                required
              ></input>

              <p id="text-header-coures">ชื่อหลักสูตร</p>
              <div className="mb-3">
                <label className="form-label" id="input-course">
                  ชื่อภาษาไทย
                </label>
                <input
                  className="form-control"
                  name="thai_course"
                  value={formData.thai_course}
                  onChange={handleChange}
                  required
                />
              </div>
              <div className="mb-3">
                <label className="form-label" id="input-course">
                  ชื่ออังกฤษ
                </label>
                <input
                  className="form-control"
                  name="eng_course"
                  value={formData.eng_course}
                  onChange={handleChange}
                />
              </div>

              <p id="text-header-coures">ชื่อปริญญา</p>
              <div className="mb-3">
                <label className="form-label" id="input-course">
                  ชื่อภาษาไทย
                </label>
                <input
                  className="form-control"
                  name="thai_degree"
                  value={formData.thai_degree}
                  onChange={handleChange}
                />
              </div>

              <div className="mb-3">
                <label className="form-label" id="input-course">
                  ชื่อภาษาอังกฤษ
                </label>
                <input
                  className="form-control"
                  name="eng_degree"
                  value={formData.eng_degree}
                  onChange={handleChange}
                />
              </div>

              {/* Textarea fields */}
              <div className="mb-3">
                <p id="text-header-coures">เกณฑ์การเข้าศึกษา</p>
                <textarea
                  className="form-control"
                  rows={3}
                  name="admission_req"
                  value={formData.admission_req}
                  onChange={handleChange}
                />
              </div>

              <div className="mb-3">
                <p id="text-header-coures">เกณฑ์สำเร็จการศึกษา</p>
                <textarea
                  className="form-control"
                  rows={3}
                  name="graduation_req"
                  value={formData.graduation_req}
                  onChange={handleChange}
                />
              </div>

              <div className="mb-3">
                <p id="text-header-coures">ปรัชญาของหลักสูตร</p>
                <textarea
                  className="form-control"
                  rows={4}
                  name="philosophy"
                  value={formData.philosophy}
                  onChange={handleChange}
                />
              </div>

              <div className="mb-3">
                <p id="text-header-coures">วัตถุประสงค์ของหลักสูตร</p>
                <textarea
                  className="form-control"
                  rows={5}
                  name="objective"
                  value={formData.objective}
                  onChange={handleChange}
                />
              </div>

              <div className="mb-3">
                <p id="text-header-coures">PLOs</p>
                <textarea
                  className="form-control"
                  rows={6}
                  name="plo"
                  value={formData.plo}
                  onChange={handleChange}
                />
              </div>

              <div className="mb-3">
                <p id="text-header-coures">อาชีพที่ประกอบได้</p>
                <textarea
                  className="form-control"
                  rows={5}
                  name="career_paths"
                  value={formData.career_paths}
                  onChange={handleChange}
                />
              </div>

              <p id="text-header-coures">ค่าใช้จ่าย</p>
              <input
                className="form-control mb-3"
                name="tuition"
                value={formData.tuition}
                onChange={handleChange}
              />

              <p id="text-header-coures">หน่วยกิต</p>
              <input
                className="form-control mb-3"
                name="credits"
                value={formData.credits}
                onChange={handleChange}
              />

              <p id="text-header-coures">รายละเอียดเพิ่มเติม</p>
              <input
                className="form-control mb-3"
                name="detail_url"
                value={formData.detail_url}
                onChange={handleChange}
              />

              <p id="text-header-coures">โครงสร้างหลักสูตร</p>
              <input
                type="file"
                className="form-control mb-3"
                name="structure"
                onChange={handleFileChange}
              />

              <p id="text-header-coures">แผนการศึกษา</p>
              <input
                type="file"
                className="form-control mb-3"
                name="roadmap"
                onChange={handleFileChange}
              />

              <button type="submit" className="btn btn-primary" id="btn-submit">
                บันทึกข้อมูลหลักสูตร
              </button>
            </form>
          </div>
        </div>
      </div>

      <Footer />
    </>
  );
};

export default Add_course;

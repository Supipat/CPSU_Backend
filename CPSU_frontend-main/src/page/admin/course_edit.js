import React, { useEffect, useState } from "react";
import Headers from "../../component/header";
import Navbar from "../../component/navbar";
import Footer from "../../component/footer";
import Menu from "../../component/menu";
import "../../css/admin/course_detail.css";
import axios from "axios";
import { useParams, useNavigate } from "react-router-dom";

const Edit_course = () => {
  const { id } = useParams();
  const [courseDetail, setCourseDetail] = useState(null);
  const [editMode, setEditMode] = useState(true);
  const [formData, setFormData] = useState({});
  const [subjects, setSubjects] = useState([]);
  const [structures, setStructures] = useState([]);
  const [roadmap, setRoadmap] = useState([]);
  const navigate = useNavigate();

  // ฟังก์ชันโหลดข้อมูลหลักสูตรและข้อมูลที่เกี่ยวข้อง
  const fetchCourseData = () => {
    axios
      .get(`http://localhost:8080/api/v1/admin/course/${id}`)
      .then((res) => {
        setCourseDetail(res.data);
        setFormData(res.data);
      })
      .catch((err) => {
        console.error("Error fetching course detail:", err);
      });

    axios
      .get(`http://localhost:8080/api/v1/admin/roadmap`)
      .then((res) => {
        // กรอง roadmap เฉพาะหลักสูตรนี้
        const courseRoadmap = res.data.filter((item) => item.course_id === id);
        setRoadmap(courseRoadmap);
      })
      .catch((err) => {
        console.error("Error fetching roadmap:", err);
      });

    axios
      .get("http://localhost:8080/api/v1/admin/subject")
      .then((res) => {
        setSubjects(res.data);
      })
      .catch((err) => {
        console.error("Error fetching subject:", err);
      });

    axios
      .get("http://localhost:8080/api/v1/admin/structure")
      .then((res) => {
        // กรอง structure เฉพาะหลักสูตรนี้
        const courseStructures = res.data.filter(
          (item) => item.course_id === id
        );
        setStructures(courseStructures);
      })
      .catch((err) => {
        console.error("Error fetching structure:", err);
      });
  };

  useEffect(() => {
    fetchCourseData();
  }, [id]);

  const saveChanges = async () => {
    try {
      await axios.put(
        `http://localhost:8080/api/v1/admin/course/${id}`,
        formData
      );
      alert("แก้ไขหลักสูตรสำเร็จ");
      navigate("/admin/course");
      setCourseDetail(formData);
      setEditMode(false);
    } catch (error) {
      console.error("Error updating course:", error);
      alert("เกิดข้อผิดพลาดในการแก้ไขข้อมูล");
    }
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const onStructureImagesChange = async (e) => {
    const files = e.target.files;
    if (!files.length || !courseDetail?.course_id) return;

    const formDataUpload = new FormData();
    formDataUpload.append("course_structure_url", files[0]);
    formDataUpload.append("thai_course", courseDetail.thai_course);
    formDataUpload.append("course_id", courseDetail.course_id);

    try {
      await axios.post(
        "http://localhost:8080/api/v1/admin/structure",
        formDataUpload,
        {
          headers: {
            "Content-Type": "multipart/form-data",
          },
        }
      );
      alert("อัปโหลดโครงสร้างหลักสูตรสำเร็จ");
      fetchCourseData(); // โหลดข้อมูลใหม่โดยไม่ต้อง reload หน้า
    } catch (error) {
      console.error("Error uploading structure:", error);
      alert("เกิดข้อผิดพลาดในการอัปโหลดโครงสร้างหลักสูตร");
    }
  };

  const onRoadmapImagesChange = async (e) => {
    const files = e.target.files;
    if (!files.length || !courseDetail?.course_id) return;

    const formDataUpload = new FormData();
    formDataUpload.append("roadmap_url", files[0]);
    formDataUpload.append("thai_course", courseDetail.thai_course);
    formDataUpload.append("course_id", courseDetail.course_id);

    try {
      await axios.post(
        "http://localhost:8080/api/v1/admin/roadmap",
        formDataUpload,
        {
          headers: {
            "Content-Type": "multipart/form-data",
          },
        }
      );
      alert("อัปโหลดแผนการเรียนสำเร็จ");
      fetchCourseData(); // โหลดข้อมูลใหม่โดยไม่ต้อง reload หน้า
    } catch (error) {
      console.error("Error uploading roadmap:", error);
      alert("เกิดข้อผิดพลาดในการอัปโหลดแผนการเรียน");
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
            <div id="group-btn-header-detail">
              <p id="news-name">
                {courseDetail?.thai_course || "กำลังโหลด..."}
              </p>
            </div>

            {/* รหัสของหลักสูตร */}
            <p id="text-header-coures">รหัสของหลักสูตร</p>
            <div id="text-content-course2">{courseDetail?.course_id}</div>
            {editMode ? (
              <textarea
                type="text"
                name="course_id"
                value={formData.course_id || ""}
                onChange={handleChange}
                className="form-control"
              />
            ) : (
              courseDetail?.course_id
            )}
            <hr />

            {/* ชื่อหลักสูตร */}
            <p id="text-header-coures">ชื่อหลักสูตร</p>
            <div className="container">
              <div className="row">
                <div className="col-4" id="text-title-course">
                  ภาษาไทย
                </div>
                <div className="col-6" id="text-content-course">
                  {editMode ? (
                    <textarea
                      type="text"
                      name="thai_course"
                      value={formData.thai_course || ""}
                      onChange={handleChange}
                      className="form-control"
                    />
                  ) : (
                    courseDetail?.thai_course
                  )}
                </div>
              </div>
              <div className="row">
                <div className="col-4" id="text-title-course">
                  ภาษาอังกฤษ
                </div>
                <div className="col-6" id="text-content-course">
                  {editMode ? (
                    <textarea
                      type="text"
                      name="eng_course"
                      value={formData.eng_course || ""}
                      onChange={handleChange}
                      className="form-control"
                    />
                  ) : (
                    courseDetail?.eng_course
                  )}
                </div>
              </div>
            </div>
            <hr />

            {/* ชื่อปริญญา */}
            <p id="text-header-coures">ชื่อปริญญา</p>
            <div className="container">
              <div className="row">
                <div className="col-4" id="text-title-course">
                  ภาษาไทย
                </div>
                <div className="col-6" id="text-content-course">
                  {editMode ? (
                    <textarea
                      type="text"
                      name="thai_degree"
                      value={formData.thai_degree || ""}
                      onChange={handleChange}
                      className="form-control"
                    />
                  ) : (
                    courseDetail?.thai_degree
                  )}
                </div>
              </div>
              <div className="row">
                <div className="col-4" id="text-title-course">
                  ภาษาอังกฤษ
                </div>
                <div className="col-6" id="text-content-course">
                  {editMode ? (
                    <textarea
                      type="text"
                      name="eng_degree"
                      value={formData.eng_degree || ""}
                      onChange={handleChange}
                      className="form-control"
                    />
                  ) : (
                    courseDetail?.eng_degree
                  )}
                </div>
              </div>
            </div>
            <hr />

            {/* เกณฑ์การเข้าศึกษาและเกณฑ์การสำเร็จการศึกษา */}
            <p id="text-header-coures">
              เกณฑ์การเข้าศึกษาและเกณฑ์การสำเร็จการศึกษา
            </p>
            <div className="container">
              <div className="row">
                <div className="col-4" id="text-title-course">
                  เกณฑ์การเข้าศึกษา
                </div>
                <div className="col-6" id="text-content-course">
                  {editMode ? (
                    <textarea
                      name="admission_req"
                      value={formData.admission_req || ""}
                      onChange={handleChange}
                      className="form-control"
                      rows={3}
                    />
                  ) : (
                    courseDetail?.admission_req
                  )}
                </div>
              </div>
              <div className="row">
                <div className="col-4" id="text-title-course">
                  เกณฑ์การสำเร็จการศึกษา
                </div>
                <div className="col-6" id="text-content-course">
                  {editMode ? (
                    <textarea
                      name="graduation_req"
                      value={formData.graduation_req || ""}
                      onChange={handleChange}
                      className="form-control"
                      rows={3}
                    />
                  ) : (
                    courseDetail?.graduation_req
                  )}
                </div>
              </div>
            </div>
            <hr />

            {/* ปรัชญาของหลักสูตร */}
            <p id="text-header-coures">ปรัชญาของหลักสูตร</p>
            <div id="text-content-course2">
              {editMode ? (
                <textarea
                  name="philosophy"
                  value={formData.philosophy || ""}
                  onChange={handleChange}
                  className="form-control"
                  rows={5}
                />
              ) : (
                courseDetail?.philosophy
              )}
            </div>
            <hr />

            {/* วัตถุประสงค์ */}
            <p id="text-header-coures">วัตถุประสงค์</p>
            <div id="text-content-course2">
              {editMode ? (
                <textarea
                  name="objective"
                  value={formData.objective || ""}
                  onChange={handleChange}
                  className="form-control"
                  rows={5}
                />
              ) : (
                courseDetail?.objective
              )}
            </div>
            <hr />

            {/* รายละเอียดผลการเรียนรู้ที่คาดหวังของหลักสูตร (PLO) */}
            <p id="text-header-coures">
              รายละเอียดผลการเรียนรู้ที่คาดหวังของหลักสูตร (PLO)
            </p>
            <div id="text-content-course2">
              {editMode ? (
                <textarea
                  name="plo"
                  value={formData.plo || ""}
                  onChange={handleChange}
                  className="form-control"
                  rows={15}
                />
              ) : (
                courseDetail?.plo
              )}
            </div>
            <hr />

            {/* ค่าใช้จ่าย */}
            <p id="text-header-coures">ค่าใช้จ่าย</p>
            <div id="text-content-course2">
              {editMode ? (
                <textarea
                  name="tuition"
                  value={formData.tuition || ""}
                  onChange={handleChange}
                  className="form-control"
                  rows={5}
                />
              ) : (
                courseDetail?.tuition
              )}
            </div>
            <hr />
            {/* หน่วยกิต */}
            <p id="text-header-coures">หน่วยกิต</p>
            <div id="text-content-course2">
              {editMode ? (
                <textarea
                  name="credits"
                  value={formData.credits || ""}
                  onChange={handleChange}
                  className="form-control"
                  rows={5}
                />
              ) : (
                courseDetail?.credits
              )}
            </div>
            <hr />

            {/* อาชีพที่สามารถประกอบได้ */}
            <p id="text-header-coures">อาชีพที่สามารถประกอบได้</p>
            <div id="text-content-course2">
              {editMode ? (
                <textarea
                  name="career_paths"
                  value={formData.career_paths || ""}
                  onChange={handleChange}
                  className="form-control"
                  rows={9}
                />
              ) : (
                courseDetail?.career_paths
              )}
            </div>
            <hr />
            {/* รายละเอียดเพิ่มเติม */}
            <p id="text-header-coures">รายละเอียดเพิ่มเติม</p>
            <div id="text-content-course2">
              {editMode ? (
                <textarea
                  name="detail_url"
                  value={formData.detail_url || ""}
                  onChange={handleChange}
                  className="form-control"
                  rows={3}
                />
              ) : (
                courseDetail?.detail_url
              )}
            </div>
            <hr />

            {/* โครงสร้างหลักสูตร */}
            <p id="text-header-coures">โครงสร้างหลักสูตร</p>

            {structures.length > 0 ? (
              structures.map((item) => (
                <div key={item.structure_id} className="mb-3">
                  <img
                    src={item.course_structure_url}
                    alt={item.thai_course}
                    style={{ maxWidth: "100%", maxHeight: "400px" }}
                  />
                </div>
              ))
            ) : (
              <p>ยังไม่มีไฟล์โครงสร้างหลักสูตร</p>
            )}

            <input
              className="form-control"
              type="file"
              onChange={onStructureImagesChange}
            />

            <hr />

            {/* แผนการเรียน (Roadmap) */}
            <p id="text-header-coures">แผนการเรียน</p>

            {roadmap.length > 0 ? (
              roadmap.map((item, index) =>
                item.roadmap_url ? (
                  <div key={index} className="mb-3">
                    <img
                      src={item.roadmap_url}
                      alt={`roadmap-${index}`}
                      style={{ maxWidth: "100%", maxHeight: "400px" }}
                    />
                  </div>
                ) : null
              )
            ) : (
              <p>ยังไม่มีไฟล์แผนการเรียน</p>
            )}

            <input
              className="form-control"
              type="file"
              onChange={onRoadmapImagesChange}
            />

            <hr />

            {/* ปุ่มบันทึกและลบ */}
            {editMode ? (
              <button className="btn btn-primary me-2" onClick={saveChanges}>
                บันทึก
              </button>
            ) : (
              <button
                className="btn btn-secondary me-2"
                onClick={() => setEditMode(true)}
              >
                บันทึก
              </button>
            )}
          </div>
        </div>
      </div>
      <Footer />
    </>
  );
};

export default Edit_course;

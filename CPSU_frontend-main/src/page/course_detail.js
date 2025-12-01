import React, { useEffect, useState } from "react";
import { useParams, Link } from "react-router-dom";
import axios from "axios";
import Navbar from "../component/navbar";
import Headers from "../component/header";
import Footer from "../component/footer";
import "../css/course_detail.css";
import { useSearchParams } from "react-router-dom";

const Course_detail = () => {
  const { course_id } = useParams();
  const [course, setCourse] = useState(null);
  const [structure, setStructure] = useState(null);
  const [roadmap, setRoadmap] = useState(null);
  const [subjects, setSubjects] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  // const [activeTab, setActiveTab] = useState("overview");

  const [searchParams] = useSearchParams();
  const initialTab = searchParams.get("tab") || "overview"; // fallback tab
  const [activeTab, setActiveTab] = useState(initialTab);

  useEffect(
    () => {
      const tabParam = searchParams.get("tab");
      if (tabParam && tabParam !== activeTab) {
        setActiveTab(tabParam);
      }
      const fetchData = async () => {
        setLoading(true);
        setError(null);

        try {
          // ดึงข้อมูลหลักสูตรทั้งหมด
          const courseRes = await axios.get(
            "http://localhost:8080/api/v1/admin/course"
          );
          const foundCourse = courseRes.data.find(
            (c) => c.course_id === course_id
          );
          setCourse(foundCourse || null);

          // ดึงข้อมูลโครงสร้างทั้งหมด
          const structureRes = await axios.get(
            "http://localhost:8080/api/v1/admin/structure"
          );
          const foundStructure = structureRes.data.find(
            (s) => s.course_id === course_id
          );
          setStructure(foundStructure || null);

          // ดึงข้อมูลแผนผังทั้งหมด
          const roadmapRes = await axios.get(
            "http://localhost:8080/api/v1/admin/roadmap"
          );
          const foundRoadmap = roadmapRes.data.find(
            (r) => r.course_id === course_id
          );
          setRoadmap(foundRoadmap || null);

          // ดึงข้อมูลรายวิชาทั้งหมด แล้วกรอง
          const subjectsRes = await axios.get(
            "http://localhost:8080/api/v1/admin/subject"
          );
          const filteredSubjects = subjectsRes.data.filter(
            (subj) => subj.course_id === course_id
          );
          setSubjects(filteredSubjects);

          setLoading(false);
        } catch (err) {
          setError("ไม่พบข้อมูลหลักสูตรที่ต้องการ");
          setLoading(false);
        }
      };

      fetchData();
    },
    [course_id],
    [searchParams]
  );

  const groupSubjectsByYearSemesterAndPlan = (subjects) => {
    const grouped = {};

    subjects.forEach((subj) => {
      const year = subj.study_year || " ";
      const semester = subj.semester || "เทอมที่ไม่ระบุ";
      const plan = subj.plan_type || "ทั่วไป"; // หรือชื่อฟิลด์จริงที่มีข้อมูลประเภทแผน

      if (!grouped[year]) grouped[year] = {};
      if (!grouped[year][semester]) grouped[year][semester] = {};
      if (!grouped[year][semester][plan]) grouped[year][semester][plan] = [];

      grouped[year][semester][plan].push(subj);
    });

    return grouped;
  };

  const groupedSubjects = groupSubjectsByYearSemesterAndPlan(subjects);

  if (loading)
    return (
      <>
        <Headers />
        <Navbar />
        <div className="container text-center my-5">
          <div className="spinner-border text-primary" role="status"></div>
          <p className="mt-3">กำลังโหลดข้อมูลหลักสูตร...</p>
        </div>
        <Footer />
      </>
    );

  if (!course)
    return (
      <>
        <Headers />
        <Navbar />
        <div className="container text-center my-5">
          <h3>ไม่พบหลักสูตรที่คุณเลือก</h3>
          <Link to="/course">กลับไปหน้ารายการหลักสูตร</Link>
        </div>
        <Footer />
      </>
    );

  return (
    <>
      <Headers />
      <Navbar />
      <div className="container my-5">
        <h2>
          <Link to="/course" id="back-to-course" className="back-button">
            &larr; กลับ
          </Link>
          {course.thai_course} ({course.year})
        </h2>
        <h4>{course.eng_course}</h4>
        <h5>{course.degree}</h5>

        {/* Tab navigation */}
        <ul className="nav nav-tabs" id="tab-course">
          <li className="nav-item">
            <button
              className={`nav-link ${activeTab === "overview" ? "active" : ""}`}
              onClick={() => setActiveTab("overview")}
            >
              ภาพรวมหลักสูตร
            </button>
          </li>
          <li className="nav-item">
            <button
              className={`nav-link ${
                activeTab === "structure" ? "active" : ""
              }`}
              onClick={() => setActiveTab("structure")}
            >
              โครงสร้างหลักสูตร
            </button>
          </li>
          <li className="nav-item">
            <button
              className={`nav-link ${activeTab === "roadmap" ? "active" : ""}`}
              onClick={() => setActiveTab("roadmap")}
            >
              แผนผังหลักสูตร
            </button>
          </li>
          <li className="nav-item">
            <button
              className={`nav-link ${activeTab === "subjects" ? "active" : ""}`}
              onClick={() => setActiveTab("subjects")}
            >
              รายวิชา
            </button>
          </li>
        </ul>

        {/* Tab content */}
        <div className="tab-content mt-3" id="course-tab-content">
          {activeTab === "overview" && (
            <div>
              <div className="mb-3">
                <h4>ชื่อหลักสูตร</h4>
                <div className="text-course-overview2">
                  <h5>ภาษาไทย</h5>
                  <p>{course.thai_course}</p>
                </div>
                <div className="text-course-overview2">
                  <h5>ภาษาอังกฤษ</h5>
                  <p>{course.eng_course}</p>
                </div>
              </div>
              <hr />
              <div className="text-course-overview">
                <h4>ชื่อปริญญา</h4>
                <div className="text-course-overview2">
                  <h5>ภาษาไทย</h5>
                  <p>{course.thai_degree}</p>
                </div>
                <div className="text-course-overview2">
                  <h5>ภาษาอังกฤษ</h5>
                  <p>{course.eng_degree}</p>
                </div>
              </div>
              <hr />
              <div className="text-overview">
                <h4>เกณท์การเข้าศึกษาและเกณท์การสำเร็จการศึกษา</h4>
                <div className="text-course-overview2">
                  <h5>เกณฑ์การเข้าศึกษา</h5>
                  <p>{course.admission_req}</p>
                </div>
                <div className="text-course-overview2">
                  <h5>เกณฑ์การสำเร็จการศึกษา</h5>
                  <p>{course.graduation_req}</p>
                </div>
              </div>
              <hr />
              <div className="mb-3">
                <h5>ปรัชญาของหลักสูตร</h5>
                <p>{course.philosophy}</p>
              </div>
              <hr />
              <div className="mb-3">
                <h5>วัตถุประสงค์</h5>
                <p>{course.objective}</p>
              </div>
              <hr />
              <div className="mb-3">
                <h5>ค่าใช้จ่าย</h5>
                <p>{course.tuition}</p>
              </div>
              <hr />
              <div className="mb-3">
                <h5>หน่วยกิต</h5>
                <p>{course.credits}</p>
              </div>
              <hr />
              <div className="mb-3">
                <h5>อาชีพที่สามารถประกอบได้หลังสําเร็จการศึกษา</h5>
                <div>
                  {course.career_paths.split("\n").map((item, index) => (
                    <p key={index}>{item}</p>
                  ))}
                </div>
              </div>
              <hr />
              <div className="mb-3">
                <h5>
                  รายละเอียดผลการเรียนรู้ที่คาดหวังของหลักสูตร (Program Learning
                  Outcome: PLOs)
                </h5>
                <div>
                  {course.plo.split("\n").map((item, index) => (
                    <p key={index}>{item}</p>
                  ))}
                </div>
              </div>
              <hr />
              <div className="mb-3">
                <h5>รายละเอียดเพิ่มเติม</h5>
                <a
                  href={course.detail_url}
                  target="_blank"
                  rel="noopener noreferrer"
                >
                  {course.detail_url}
                </a>
              </div>
            </div>
          )}

          {activeTab === "structure" && (
            <div>
              <h4 id="structure-tab-name">โครงสร้างหลักสูตร</h4>
              <p id="structure-tab-name">{course.thai_course}</p>
              <p id="structure-tab-name">{course.eng_degree}</p>
              {structure ? (
                <img
                  src={structure.course_structure_url}
                  alt={`โครงสร้างหลักสูตร ${course.course_id}`}
                  className="img-fluid"
                />
              ) : (
                <p>ไม่พบข้อมูลโครงสร้างหลักสูตร</p>
              )}
            </div>
          )}

          {activeTab === "roadmap" && (
            <div>
              <h4 id="roadmap-tab-name">แผนผังหลักสูตร</h4>
              <p id="roadmap-tab-name">{course.thai_course}</p>
              <p id="roadmap-tab-name">{course.eng_degree}</p>
              {roadmap ? (
                <img
                  src={roadmap.roadmap_url}
                  alt={`แผนผังหลักสูตร ${course.course_id}`}
                  className="img-fluid"
                />
              ) : (
                <p>ไม่พบข้อมูลแผนผังหลักสูตร</p>
              )}
            </div>
          )}

          {activeTab === "subjects" && (
            <div>
              <h4>รายวิชาในหลักสูตร</h4>
              {subjects.length === 0 ? (
                <p>ไม่พบรายวิชาในหลักสูตรนี้</p>
              ) : (
                Object.entries(groupedSubjects).map(([year, semesters]) => (
                  <div key={year} style={{ marginBottom: "2rem" }}>
                    <h5>{year}</h5>
                    {Object.entries(semesters).map(([semester, plans]) => (
                      <div key={semester} style={{ marginBottom: "1rem" }}>
                        <h6>{semester}</h6>
                        {Object.entries(plans).map(([plan, subs]) => (
                          <div key={plan} style={{ marginBottom: "1rem" }}>
                            <h6 style={{ fontWeight: "bold" }}>{plan}</h6>
                            <table className="table table-striped">
                              <thead>
                                <tr>
                                  <th>รหัสรายวิชา</th>
                                  <th>ชื่อวิชา</th>
                                  <th>หน่วยกิต</th>
                                </tr>
                              </thead>
                              <tbody>
                                {subs.map((subj) => {
                                  return (
                                    <tr key={subj.subject_id}>
                                      <td>
                                        {subj.subject_id == "SUxxx" ||
                                        subj.subject_id == "SUXXX" ||
                                        subj.subject_id == "------" ? (
                                          subj.subject_id
                                        ) : (
                                          <Link to={`/subject/${subj.id}`}>
                                            {subj.subject_id}
                                          </Link>
                                        )}
                                      </td>
                                      <td>{subj.thai_subject}</td>
                                      <td>{subj.credits}</td>
                                    </tr>
                                  );
                                })}
                              </tbody>
                            </table>
                          </div>
                        ))}
                      </div>
                    ))}
                  </div>
                ))
              )}
            </div>
          )}
        </div>
      </div>
      <Footer />
    </>
  );
};

export default Course_detail;

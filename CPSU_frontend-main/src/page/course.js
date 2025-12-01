import React, { useEffect, useState } from "react";
import axios from "axios";
import Navbar from "../component/navbar";
import Headers from "../component/header";
import Footer from "../component/footer";
import "../css/course.css";
import { Link } from "react-router-dom";

const MOCK_DATA = [
  // ...ข้อมูล MOCK_DATA เหมือนเดิม
];

const ALLOWED_MAJORS = {
  ปริญญาตรี: ["วิทยาการคอมพิวเตอร์", "เทคโนโลยีสารสนเทศ", "วิทยาการข้อมูล"],
  ปริญญาโท: ["เทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล"],
  ปริญญาเอก: ["เทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล"],
};
const DEGREE_ORDER = ["ปริญญาตรี", "ปริญญาโท", "ปริญญาเอก"];

const Course = () => {
  const [courses, setCourses] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchCourses = async () => {
      try {
        const res = await axios.get(
          "http://localhost:8080/api/v1/admin/course"
        );
        if (Array.isArray(res.data) && res.data.length > 0) {
          setCourses(res.data);
        } else {
          setCourses(MOCK_DATA);
        }
      } catch (err) {
        setCourses(MOCK_DATA);
      }
      setLoading(false);
    };
    fetchCourses();
  }, []);

  // กรองข้อมูลเฉพาะสาขาที่กำหนดไว้เท่านั้น
  const filteredCourses = courses.filter(
    (course) =>
      ALLOWED_MAJORS[course.degree] &&
      ALLOWED_MAJORS[course.degree].includes(course.major)
  );

  // เรียงปีใหม่ไปเก่า
  filteredCourses.sort((a, b) => b.year - a.year);

  // จัดกลุ่มตาม degree และ major
  const grouped = {};
  filteredCourses.forEach((course) => {
    if (!grouped[course.degree]) grouped[course.degree] = {};
    if (!grouped[course.degree][course.major])
      grouped[course.degree][course.major] = [];
    grouped[course.degree][course.major].push(course);
  });

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

  return (
    <>
      <Headers />
      <Navbar />

      <div className="container my-5">
        <h2 className="mb-4 text-center">หลักสูตรทั้งหมด</h2>

        {DEGREE_ORDER.map((degree) =>
          grouped[degree] ? (
            <div key={degree} className="degree-section mb-4">
              <h3 className="degree-header mb-3">{degree}</h3>
              {Object.entries(grouped[degree]).map(([major, courseList]) => (
                <div key={major} className="major-section mb-3">
                  <h4>{major}</h4>
                  <div className="course-card-container">
                    {courseList.map((course) => (
                      <Link
                        key={course.course_id}
                        to={`/course/${course.course_id}`}
                        className="card course-card text-decoration-none text-dark"
                      >
                        <div className="card-body">
                          <h5 className="card-title">
                            {course.thai_course} ({course.year})
                          </h5>
                          <p className="card-text">{course.eng_course}</p>
                        </div>
                      </Link>
                    ))}
                  </div>
                </div>
              ))}
            </div>
          ) : null
        )}
      </div>

      <Footer />
    </>
  );
};

export default Course;

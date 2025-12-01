import React, { useEffect, useState } from "react";
import "../../css/admin/course.css";
import Headers from "../../component/header";
import Navbar from "../../component/navbar";
import Footer from "../../component/footer";
import Menu from "../../component/menu";
import { Link } from "react-router-dom";
import axios from "axios";

const Course = () => {
  const [courses, setCourses] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    axios
      .get("http://localhost:8080/api/v1/admin/course")
      .then((res) => {
        setCourses(res.data);
        setLoading(false);
      })
      .catch((err) => {
        console.error("Error fetching courses:", err);
        setError("ไม่สามารถโหลดข้อมูลหลักสูตรได้");
        setLoading(false);
      });
  }, []);

  const bachelorCourses = courses.filter((c) => c.degree === "ปริญญาตรี");
  const masterCourses = courses.filter((c) => c.degree === "ปริญญาโท");
  const doctorateCourses = courses.filter((c) => c.degree === "ปริญญาเอก");

  const renderCourseList = (courseList) => {
    if (courseList.length === 0)
      return <p className="text-muted">ไม่มีหลักสูตรในระดับนี้</p>;

    return courseList.map((course) => (
      <div className="col-md-4 mb-4" key={course.course_id}>
        <div className="card h-100 shadow-sm">
          <div className="card-body d-flex flex-column">
            <h5 className="card-title">{course.course_id}</h5>
            <h6 className="card-subtitle mb-2 text-muted">
              {course.thai_course}
            </h6>
            <div className="mt-auto">
              <Link
                to={`/admin/course/${course.course_id}`}
                className="btn btn-primary"
                id="btn-detailcourse"
              >
                ดูรายละเอียด
              </Link>
            </div>
          </div>
        </div>
      </div>
    ));
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

          <div className="col-sm-9">
            <div className="d-flex justify-content-between align-items-center mb-4">
              <h3 id="course-title">หลักสูตรทั้งหมด</h3>
              <Link to="/admin/addcourse" className="btn-addcourse">
                + เพิ่มหลักสูตร
              </Link>
            </div>

            {loading && <p>กำลังโหลดข้อมูลหลักสูตร...</p>}
            {error && <p className="text-danger">{error}</p>}

            {!loading && !error && (
              <>
                <section className="mb-5">
                  <div className="degree-divider">
                    <span className="bachelor">ปริญญาตรี</span>
                  </div>
                  <div className="row">{renderCourseList(bachelorCourses)}</div>
                </section>

                <section className="mb-5">
                  <div className="degree-divider">
                    <span className="master">ปริญญาโท</span>
                  </div>
                  <div className="row">{renderCourseList(masterCourses)}</div>
                </section>

                <section className="mb-5">
                  <div className="degree-divider">
                    <span className="doctor">ปริญญาเอก</span>
                  </div>
                  <div className="row">
                    {renderCourseList(doctorateCourses)}
                  </div>
                </section>
              </>
            )}
          </div>
        </div>
      </div>

      <Footer />
    </>
  );
};

export default Course;

import React, { useEffect } from "react";
import Navbar from "../component/navbar";
import Headers from "../component/header";
import Footer from "../component/footer";
import "../css/about.css";

const About = () => {
  useEffect(() => {
    const elements = document.querySelectorAll(".fade-in");

    const observer = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          if (entry.isIntersecting) {
            entry.target.classList.add("show");
          }
        });
      },
      {
        threshold: 0.1,
      }
    );

    elements.forEach((el) => observer.observe(el));

    return () => {
      elements.forEach((el) => observer.unobserve(el));
    };
  }, []);

  return (
    <>
      <Headers />
      <Navbar />

      <div className="container about-container">
        <div className="row justify-content-center">
          <div className="col-md-10">
            <h2 className="section-title">เกี่ยวกับภาควิชาคอมพิวเตอร์</h2>

            {/* Section: ประวัติภาควิชา */}
            <section className="about-section">
              <h4 className="text-start mb-5">ประวัติภาควิชา</h4>

              <div className="timeline white-box">
                <div className="timeline-item fade-in">
                  <div className="timeline-content">
                    <h5>พ.ศ. 2531</h5>
                    <p>
                      เริ่มเปิดสอนหลักสูตรวิทยาศาสตรบัณฑิต
                      สาขาวิทยาการคอมพิวเตอร์ เป็นครั้งแรก
                      ซึ่งยังสังกัดอยู่ในภาควิชาคณิตศาสตร์
                    </p>
                  </div>
                </div>

                <div className="timeline-item fade-in">
                  <div className="timeline-content">
                    <h5>พ.ศ. 2545</h5>
                    <p>จัดตั้งเป็นภาควิชาคอมพิวเตอร์ สังกัดคณะวิทยาศาสตร์</p>
                  </div>
                </div>

                <div className="timeline-item fade-in">
                  <div className="timeline-content">
                    <h5>พ.ศ. 2547</h5>
                    <p>
                      ได้รับการอนุมัติจากสภามหาวิทยาลัยศิลปากร ให้จัดตั้งเป็น
                      ภาควิชาคอมพิวเตอร์ สังกัดคณะวิทยาศาสตร์
                      และเริ่มเปิดสอนหลักสูตรวิทยาศาสตรมหาบัณฑิต
                      สาขาเทคโนโลยีสารสนเทศ
                    </p>
                  </div>
                </div>

                <div className="timeline-item fade-in">
                  <div className="timeline-content">
                    <h5>พ.ศ. 2548</h5>
                    <p>
                      เริ่มเปิดสอนหลักสูตรวิทยาศาสตรบัณฑิต สาขาเทคโนโลยีสารสนเทศ
                    </p>
                  </div>
                </div>
              </div>
            </section>

            {/* Section: ติดต่อภาควิชา */}
            <section className="about-section white-box">
              <h4>ติดต่อภาควิชาคอมพิวเตอร์</h4>
              <p>
                <strong>ที่อยู่ (ภาษาไทย)</strong>
                <br />
                ภาควิชาคอมพิวเตอร์ คณะวิทยาศาสตร์, มหาวิทยาลัยศิลปากร
                <br />
                วิทยาเขตพระราชวังสนามจันทร์ ถนนราชมรรคาใน
                <br />
                อำเภอเมืองฯ จังหวัดนครปฐม 73000
              </p>
              <p>
                <strong>Address (English)</strong>
                <br />
                Department of Computing, Faculty of Science, Silpakorn
                University
                <br />
                Sanam Chandra Palace Campus, Rajamunka-nai Rd.
                <br />
                Muang District, Nakhon Pathom 73000
              </p>
              <p>
                <strong>โทรศัพท์ (เบอร์ตรงภาควิชา)</strong>
                <br />
                0-3414-7014, 0-3414-7015
              </p>

              <p>
                <strong>เบอร์กลางคณะฯ</strong>
                <br />
                0-3414-7001
              </p>

              <p>
                <strong>แผนที่:</strong>
                <br />
                <div className="map-container">
                  <iframe
                    src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d1937.1693423775325!2d100.03880173870269!3d13.81869014664506!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x30fcffa1b2043d1b%3A0x18c800b2e15effc7!2sSilpakorn%20University%2C%20Sanam%20Chandra%20Palace%20Campus!5e0!3m2!1sen!2sth!4v1758449318289!5m2!1sen!2sth"
                    width="100%"
                    height="350"
                    style={{ border: 0 }}
                    allowFullScreen=""
                    loading="lazy"
                    referrerPolicy="no-referrer-when-downgrade"
                    title="แผนที่ภาควิชาคอมพิวเตอร์"
                  ></iframe>
                </div>
              </p>
            </section>
          </div>
        </div>
      </div>

      <Footer />
    </>
  );
};

export default About;

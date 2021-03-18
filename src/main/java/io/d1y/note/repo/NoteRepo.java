package io.d1y.note.repo;

import io.d1y.note.dao.NoteDao;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.transaction.annotation.Transactional;

@Transactional // 事务规范👻
public interface NoteRepo extends JpaRepository<NoteDao, Integer> {

  @Query(value = "SELECT * FROM note_dao WHERE router=?", nativeQuery = true)
  public NoteDao findRouter(String router);

  @Modifying
  @Query(value = "UPDATE note_dao SET content=:content  WHERE router=:router", nativeQuery = true)
  public void updateContent(@Param("content") String content, @Param("router") String router);

}
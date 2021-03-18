package io.d1y.note.dao;

import io.d1y.note.utils.BaseEntity;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;

@Entity
@Getter
@Setter
@EqualsAndHashCode(callSuper = true)
@Data
public class NoteDao extends BaseEntity {

  @Id
  @GeneratedValue(strategy=GenerationType.AUTO)
  private Integer id;

  private String title;

  private String router;

  private String content;

}
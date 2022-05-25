-- //////////////////////////////////////////////////////////////////////////
-- LIKE TRIGGER
-- /////////////////////////////////////////////////////////////////////////

CREATE OR REPLACE FUNCTION like_post() RETURNS TRIGGER LANGUAGE PLPGSQL AS
$$
  BEGIN
    UPDATE post SET post_like_count = post_like_count + 1,
    updated_at = NOW()
    WHERE id = NEW.post_id;
    RETURN NEW;
  END;
$$;

CREATE TRIGGER like_post_trigger AFTER INSERT ON liked_post
FOR EACH ROW EXECUTE PROCEDURE like_post();

-- //////////////////////////////////////////////////////////////////////////
-- UNLIKE TRIGGER
-- /////////////////////////////////////////////////////////////////////////

CREATE OR REPLACE FUNCTION unlike_post() RETURNS TRIGGER LANGUAGE PLPGSQL AS
$$
  BEGIN
    UPDATE post SET post_like_count = post_like_count - 1,
    updated_at = NOW()
    WHERE id = NEW.post_id;
    RETURN NEW;
  END;
$$;

CREATE TRIGGER unlike_post_trigger AFTER UPDATE ON liked_post
FOR EACH ROW EXECUTE PROCEDURE unlike_post();


-- //////////////////////////////////////////////////////////////////////////
-- VIEW TRIGGER
-- /////////////////////////////////////////////////////////////////////////

CREATE OR REPLACE FUNCTION view_post() RETURNS TRIGGER LANGUAGE PLPGSQL AS
$$
  BEGIN
    UPDATE post SET post_views_count = post_views_count + 1,
    updated_at = NOW()
    WHERE id = NEW.post_id;
    RETURN NEW;
  END;
$$;

CREATE TRIGGER view_post_trigger AFTER INSERT ON viewed_post
FOR EACH ROW EXECUTE PROCEDURE view_post();

-- //////////////////////////////////////////////////////////////////////////
-- RATING TRIGGER
-- /////////////////////////////////////////////////////////////////////////

CREATE OR REPLACE FUNCTION rating_post() RETURNS TRIGGER LANGUAGE PLPGSQL AS
$$
  BEGIN
    UPDATE post SET post_vote_count = post_vote_count + 1,
    updated_at = NOW()
    WHERE id = NEW.post_id;
    RETURN NEW;
  END;
$$;

CREATE TRIGGER rating_post_trigger AFTER INSERT ON rating_post
FOR EACH ROW EXECUTE PROCEDURE rating_post();


-- //////////////////////////////////////////////////////////////////////////
--  LIKE FUNCTION
-- /////////////////////////////////////////////////////////////////////////

CREATE OR REPLACE  FUNCTION toggle_comment_like(user_id INTEGER,like_post_id INTEGER) RETURNS VOID LANGUAGE PLPGSQL AS
$$
  BEGIN
IF  EXISTS (SELECT id  FROM liked_post  WHERE reader_id =user_id AND post_id=like_post_id AND deleted_at IS NULL)
THEN
UPDATE  liked_post SET deleted_at = NOW() WHERE reader_id = user_id AND post_id = like_post_id  ;
   ELSE
   INSERT INTO liked_post (reader_id  , post_id ) VALUES (user_id  , like_post_id)  ;
   END IF;
  END
$$;
